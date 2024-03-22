-- Filename: add_health_insurance_table.sql

-- +goose Up
CREATE TABLE IF NOT EXISTS health_insurances (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO health_insurances (name) VALUES ('Aetna'), ('HealthCare'), ('Cigna');

-- +goose Down
DROP TABLE IF EXISTS health_insurances;
