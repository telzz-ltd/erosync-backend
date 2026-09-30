-- +goose Up
CREATE table otps (
    recipient VARCHAR(255) NOT NULL, 
    code_hash TEXT NOT NULL, 
    purpose VARCHAR(50) NOT NULL,
    channel VARCHAR(50) NOT NULL,
    attempts INTEGER NOT NULL, 
    max_attempts INTEGER NOT NULL, 
    created_at TIMESTAMP DEFAULT NOW(), 
    expires_at TIMESTAMP NOT NULL,

    PRIMARY KEY (recipient, purpose, channel)
);

-- +goose Down
DROP TABLE otps;
