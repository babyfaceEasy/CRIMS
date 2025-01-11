-- +goose Up
CREATE TABLE IF NOT EXISTS customer_cloud_resources (
    customer_id INT NOT NULL,
    cloud_resource_id INT NOT NULL,
    PRIMARY KEY (customer_id, cloud_resource_id),
    CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE CASCADE,
    CONSTRAINT fk_cloud_resource FOREIGN KEY (cloud_resource_id) REFERENCES cloud_resources (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS customer_cloud_resources;
