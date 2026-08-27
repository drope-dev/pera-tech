package catalog

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct{ database *pgxpool.Pool }

func NewPostgresRepository(database *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{database: database}
}
func (r *PostgresRepository) Create(ctx context.Context, input CreateProductInput) (Product, error) {
	return scan(r.database.QueryRow(ctx, `INSERT INTO products (slug,name,description,price_cents,stock_on_hand,status) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id::text,slug,name,description,price_cents,stock_on_hand,status,created_at,updated_at`, input.Slug, input.Name, input.Description, input.PriceCents, input.StockOnHand, input.Status))
}
func (r *PostgresRepository) ListPublished(ctx context.Context) ([]Product, error) {
	rows, err := r.database.Query(ctx, `SELECT id::text,slug,name,description,price_cents,stock_on_hand,status,created_at,updated_at FROM products WHERE status='published' AND stock_on_hand>0 ORDER BY created_at DESC`)
	if err != nil {
		return nil, fmt.Errorf("list products: %w", err)
	}
	defer rows.Close()
	result := []Product{}
	for rows.Next() {
		p, e := scan(rows)
		if e != nil {
			return nil, e
		}
		result = append(result, p)
	}
	return result, rows.Err()
}
func (r *PostgresRepository) FindPublishedBySlug(ctx context.Context, slug string) (Product, error) {
	return scan(r.database.QueryRow(ctx, `SELECT id::text,slug,name,description,price_cents,stock_on_hand,status,created_at,updated_at FROM products WHERE slug=$1 AND status='published' AND stock_on_hand>0`, slug))
}

type scanner interface{ Scan(...any) error }

func scan(row scanner) (Product, error) {
	var p Product
	err := row.Scan(&p.ID, &p.Slug, &p.Name, &p.Description, &p.PriceCents, &p.StockOnHand, &p.Status, &p.CreatedAt, &p.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return Product{}, ErrNotFound
	}
	if err != nil {
		return Product{}, fmt.Errorf("scan product: %w", err)
	}
	return p, nil
}
