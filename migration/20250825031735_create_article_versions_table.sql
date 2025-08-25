-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS article_versions (
                                  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                  article_id UUID NOT NULL REFERENCES articles(id) ON DELETE CASCADE,
                                  version_number INT NOT NULL,
                                  content TEXT NOT NULL,
                                  trending_score FLOAT DEFAULT 0.0,
                                  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                                  UNIQUE(article_id, version_number)
);

CREATE INDEX idx_article_versions_article_id ON article_versions(article_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_article_versions_article_id;
DROP TABLE IF EXISTS article_versions;
-- +goose StatementEnd
