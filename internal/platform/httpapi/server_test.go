package httpapi

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drope-dev/pera-tech/internal/catalog"
)

type fakePinger struct{ err error }

func (p fakePinger) Ping(context.Context) error { return p.err }

type fakeCatalog struct{}

func (fakeCatalog) Create(context.Context, catalog.CreateProductInput) (catalog.Product, error) {
	return catalog.Product{}, nil
}
func (fakeCatalog) ListPublished(context.Context) ([]catalog.Product, error) { return nil, nil }
func (fakeCatalog) FindPublishedBySlug(context.Context, string) (catalog.Product, error) {
	return catalog.Product{}, catalog.ErrNotFound
}
func handler(pinger Pinger) http.Handler {
	return NewHandler(pinger, catalog.NewService(fakeCatalog{}), "a-valid-local-admin-token-with-32-chars")
}

func TestLiveEndpoint(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/health/live", nil)
	response := httptest.NewRecorder()

	handler(fakePinger{}).ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}
}

func TestReadyEndpointReportsDatabaseFailureWithoutDetails(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/health/ready", nil)
	response := httptest.NewRecorder()

	handler(fakePinger{err: errors.New("connection refused postgres://secret")}).ServeHTTP(response, request)

	if response.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusServiceUnavailable)
	}
	if response.Body.String() != "{\"status\":\"unavailable\"}\n" {
		t.Fatalf("body = %q, want generic unavailable response", response.Body.String())
	}
}
