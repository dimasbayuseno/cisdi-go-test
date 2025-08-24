-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS examples
(
    id
    UUID
    PRIMARY
    KEY
    DEFAULT
    gen_random_uuid
(
),
    name VARCHAR
(
    255
) NOT NULL,
    description TEXT,
    type VARCHAR
(
    255
) NOT NULL,
    created_at TIMESTAMP DEFAULT now
(
),
    updated_at TIMESTAMP DEFAULT now
(
)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS examples;
-- +goose StatementEnd
