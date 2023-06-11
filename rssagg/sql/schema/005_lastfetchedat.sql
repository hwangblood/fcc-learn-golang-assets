-- +goose Up
ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP;  -- default to null

-- +goose Down
ALTER TABLE feeds DROP COLUMN last_fetched_at;