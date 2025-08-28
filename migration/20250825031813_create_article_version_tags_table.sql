-- +goose Up
-- +goose StatementBegin
CREATE TABLE article_version_tags (
                                      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                      article_version_id UUID NOT NULL REFERENCES article_versions(id) ON DELETE CASCADE,
                                      tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
                                      article_tag_relationship_score FLOAT DEFAULT 0.0,
                                      created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                                      updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                                      UNIQUE(article_version_id, tag_id)
);

CREATE INDEX idx_article_version_tags_article_version_id ON article_version_tags(article_version_id);
CREATE INDEX idx_article_version_tags_tag_id ON article_version_tags(tag_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_article_version_tags_tag_id;
DROP INDEX IF EXISTS idx_article_version_tags_article_version_id;
DROP TABLE IF EXISTS article_version_tags;
-- +goose StatementEnd
