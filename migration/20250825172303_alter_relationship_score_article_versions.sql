-- +goose Up
-- +goose StatementBegin
ALTER TABLE article_versions
    ADD COLUMN article_tag_relationship_score FLOAT DEFAULT 0.0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE article_versions
    DROP COLUMN article_tag_relationship_score;
-- +goose StatementEnd
