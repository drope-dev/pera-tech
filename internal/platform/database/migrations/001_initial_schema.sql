-- Loop 001 will introduce the first business tables.
-- This migration establishes an explicit, versioned schema baseline.
CREATE TABLE application_metadata (
    key TEXT PRIMARY KEY,
    value TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
