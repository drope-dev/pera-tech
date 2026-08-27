CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug TEXT NOT NULL UNIQUE CHECK (slug ~ '^[a-z0-9]+(-[a-z0-9]+)*$'),
    name TEXT NOT NULL CHECK (char_length(name) BETWEEN 1 AND 160),
    description TEXT NOT NULL DEFAULT '' CHECK (char_length(description) <= 5000),
    price_cents BIGINT NOT NULL CHECK (price_cents BETWEEN 1 AND 100000000),
    stock_on_hand INTEGER NOT NULL CHECK (stock_on_hand BETWEEN 0 AND 1000000),
    status TEXT NOT NULL CHECK (status IN ('draft', 'published', 'archived')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(), updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX products_public_catalog_idx ON products (created_at DESC) WHERE status = 'published' AND stock_on_hand > 0;
