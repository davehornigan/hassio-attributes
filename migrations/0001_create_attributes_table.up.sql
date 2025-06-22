CREATE TABLE attributes
(
    id             SERIAL PRIMARY KEY,
    user_id        VARCHAR(64) NOT NULL,
    attribute_name VARCHAR(64) NOT NULL,
    json_value     JSONB       NOT NULL,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);