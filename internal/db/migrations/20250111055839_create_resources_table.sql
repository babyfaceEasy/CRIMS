-- +goose Up
CREATE TABLE IF NOT EXISTS cloud_resources (
    id bigserial NOT NULL PRIMARY KEY,
    uid UUID DEFAULT uuid_generate_v4(), -- Generates a unique UUID V4 for each customer
    customer_id bigint,
    name varchar(200) UNIQUE NOT NULL,
    type varchar(100),
    region varchar(100),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);

-- +goose Down
DROP TABLE IF EXISTS cloud_resources;