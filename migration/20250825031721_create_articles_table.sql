-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS articles (
                          id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          author_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                          title VARCHAR(255) NOT NULL,
                          slug VARCHAR(255) UNIQUE NOT NULL,
                          status VARCHAR(50) NOT NULL CHECK (status IN ('draft', 'published', 'archived')) DEFAULT 'draft',
                          created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                          updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                          published_at TIMESTAMPTZ
);

CREATE INDEX idx_articles_author_id ON articles(author_id);
CREATE INDEX idx_articles_status ON articles(status);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_articles_status;
DROP INDEX IF EXISTS idx_articles_author_id;
DROP TABLE IF EXISTS articles;
-- +goose StatementEnd
