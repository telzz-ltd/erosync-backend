-- +goose Up
CREATE TABLE brands(
    id TEXT NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description VARCHAR(255),
    logo_url TEXT,
    contact_info JSONB DEFAULT '{}',
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE brand_categories (
    id TEXT NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description VARCHAR(255)
);

CREATE TABLE brand_category_pivot (
    brand_id TEXT NOT NULL REFERENCES brands(id),
    category_id TEXT NOT NULL REFERENCES brand_categories(id),
    
    PRIMARY KEY (brand_id, category_id)
);

-- +goose Down
DROP TABLE brands;
DROP TABLE brand_categories;
DROP TABLE brand_category_pivot;
