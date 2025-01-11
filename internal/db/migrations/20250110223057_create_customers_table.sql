-- +goose Up
CREATE TABLE IF NOT EXISTS customers (
    id bigserial NOT NULL PRIMARY KEY,
    uid UUID DEFAULT uuid_generate_v4(), -- Generates a unique UUID V4 for each customer
    name varchar(100),
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone 
);

-- +goose Down
DROP TABLE IF EXISTS customers;
