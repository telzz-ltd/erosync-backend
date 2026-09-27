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
    
    CONSTRAINT pk_brand_category_pivot PRIMARY KEY (brand_id, category_id),

    CONSTRAINT fk_brand_category_pivot_brand_id 
        FOREIGN KEY (brand_id) 
        REFERENCES brands(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_brand_category_pivot_category_id 
        FOREIGN KEY (category_id) 
        REFERENCES brand_categories(id) 
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE brand_category_pivot;
DROP TABLE brand_categories;
DROP TABLE brands;
