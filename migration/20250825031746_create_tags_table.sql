-- +goose Up
-- +goose StatementBegin
CREATE TABLE tags (
                      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                      name VARCHAR(100) UNIQUE NOT NULL,
                      usage_count INT DEFAULT 0,
                      last_used_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                      created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tags;
-- +goose StatementEnd
