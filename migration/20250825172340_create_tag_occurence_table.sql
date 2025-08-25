-- +goose Up
-- +goose StatementBegin
CREATE TABLE tag_cooccurrence (
                                  tag_a_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
                                  tag_b_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
                                  cooccurrence_count INT NOT NULL DEFAULT 0,
                                  PRIMARY KEY (tag_a_id, tag_b_id)
);
CREATE INDEX idx_tag_cooccurrence_a ON tag_cooccurrence(tag_a_id);
CREATE INDEX idx_tag_cooccurrence_b ON tag_cooccurrence(tag_b_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_tag_cooccurrence_b;
DROP INDEX IF EXISTS idx_tag_cooccurrence_a;
DROP TABLE IF EXISTS tag_cooccurrence;
-- +goose StatementEnd
