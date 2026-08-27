package httpapi

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/drope-dev/pera-tech/internal/catalog"
)

type Pinger interface {
	Ping(context.Context) error
}

func NewHandler(database Pinger, products *catalog.Service, adminToken string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health/live", live)
	mux.HandleFunc("GET /health/ready", ready(database))
	mux.HandleFunc("GET /v1/products", func(w http.ResponseWriter, r *http.Request) {
		products, err := products.ListPublished(r.Context())
		if err != nil {
			writeJSON(w, 500, map[string]string{"error": "internal error"})
			return
		}
		writeValue(w, 200, products)
	})
	mux.HandleFunc("GET /v1/products/{slug}", func(w http.ResponseWriter, r *http.Request) {
		product, err := products.FindPublishedBySlug(r.Context(), r.PathValue("slug"))
		if errors.Is(err, catalog.ErrNotFound) {
			writeJSON(w, 404, map[string]string{"error": "not found"})
			return
		}
		if err != nil {
			writeJSON(w, 500, map[string]string{"error": "internal error"})
			return
		}
		writeValue(w, 200, product)
	})
	mux.HandleFunc("POST /v1/admin/products", func(w http.ResponseWriter, r *http.Request) {
		if !authorized(r.Header.Get("Authorization"), adminToken) {
			writeJSON(w, 401, map[string]string{"error": "unauthorized"})
			return
		}
		var input catalog.CreateProductInput
		decoder := json.NewDecoder(io.LimitReader(r.Body, 1<<20))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&input); err != nil {
			writeJSON(w, 400, map[string]string{"error": "invalid request"})
			return
		}
		product, err := products.Create(r.Context(), input)
		if errors.Is(err, catalog.ErrInvalidInput) {
			writeJSON(w, 400, map[string]string{"error": "invalid request"})
			return
		}
		if err != nil {
			writeJSON(w, 500, map[string]string{"error": "internal error"})
			return
		}
		writeValue(w, 201, product)
	})
	return mux
}

func authorized(header, expected string) bool {
	const prefix = "Bearer "
	if !strings.HasPrefix(header, prefix) {
		return false
	}
	provided := strings.TrimPrefix(header, prefix)
	return len(provided) == len(expected) && subtle.ConstantTimeCompare([]byte(provided), []byte(expected)) == 1
}
func writeValue(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func live(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func ready(database Pinger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		if err := database.Ping(ctx); err != nil {
			writeJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "unavailable"})
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}

func writeJSON(w http.ResponseWriter, status int, body map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
