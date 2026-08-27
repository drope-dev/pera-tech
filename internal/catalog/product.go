package catalog

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

var slugPattern = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

var (
	ErrNotFound     = errors.New("product not found")
	ErrInvalidInput = errors.New("invalid product input")
)

type ProductStatus string

const (
	StatusDraft     ProductStatus = "draft"
	StatusPublished ProductStatus = "published"
	StatusArchived  ProductStatus = "archived"
)

type Product struct {
	ID          string        `json:"id"`
	Slug        string        `json:"slug"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	PriceCents  int64         `json:"price_cents"`
	StockOnHand int           `json:"stock_on_hand"`
	Status      ProductStatus `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type CreateProductInput struct {
	Slug        string        `json:"slug"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	PriceCents  int64         `json:"price_cents"`
	StockOnHand int           `json:"stock_on_hand"`
	Status      ProductStatus `json:"status"`
}

type Repository interface {
	Create(context.Context, CreateProductInput) (Product, error)
	ListPublished(context.Context) ([]Product, error)
	FindPublishedBySlug(context.Context, string) (Product, error)
}

type Service struct{ repository Repository }

func NewService(repository Repository) *Service { return &Service{repository: repository} }
func (s *Service) Create(ctx context.Context, input CreateProductInput) (Product, error) {
	input.Slug = strings.ToLower(strings.TrimSpace(input.Slug))
	input.Name = strings.TrimSpace(input.Name)
	input.Description = strings.TrimSpace(input.Description)
	if !slugPattern.MatchString(input.Slug) || len(input.Slug) > 120 || input.Name == "" || len(input.Name) > 160 || len(input.Description) > 5000 || input.PriceCents < 1 || input.PriceCents > 100_000_000 || input.StockOnHand < 0 || input.StockOnHand > 1_000_000 || (input.Status != StatusDraft && input.Status != StatusPublished && input.Status != StatusArchived) {
		return Product{}, fmt.Errorf("%w", ErrInvalidInput)
	}
	return s.repository.Create(ctx, input)
}
func (s *Service) ListPublished(ctx context.Context) ([]Product, error) {
	return s.repository.ListPublished(ctx)
}
func (s *Service) FindPublishedBySlug(ctx context.Context, slug string) (Product, error) {
	slug = strings.TrimSpace(slug)
	if !slugPattern.MatchString(slug) {
		return Product{}, ErrNotFound
	}
	return s.repository.FindPublishedBySlug(ctx, slug)
}
