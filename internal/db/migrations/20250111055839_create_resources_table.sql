-- +goose Up
CREATE TABLE IF NOT EXISTS resources (
    id bigserial NOT NULL PRIMARY KEY,
    uid UUID DEFAULT uuid_generate_v4(), -- Generates a unique UUID V4 for each customer
    customer_id bigint,
    name varchar(200) UNIQUE NOT NULL,
    type varchar(50),
    region varchar(50),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT fk_customers_resources FOREIGN KEY(customer_id) REFERENCES customers(id) ON UPDATE CASCADE ON DELETE CASCADE,
);

-- +goose Down
DROP TABLE IF EXISTS resources;