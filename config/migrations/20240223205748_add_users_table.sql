-- +goose Up
CREATE TYPE gender_type AS ENUM ('male', 'female', 'other');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INT,
    gender gender_type NOT NULL,
    phone_number VARCHAR(20)
);

-- Add a check constraint to enforce the gender_type values
ALTER TABLE users
ADD CONSTRAINT gender_type_check
CHECK (gender IN ('male', 'female', 'other'));

-- +goose Down
DROP TABLE users;

-- Drop the custom ENUM type
DROP TYPE gender_type;
