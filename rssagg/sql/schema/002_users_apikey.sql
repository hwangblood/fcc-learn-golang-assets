-- +goose Up
ALTER TABLE users ADD COLUMN api_key varchar(64) UNIQUE NOT NULL DEFAULT (
    -- generate some hashed sha256 text, and cast it into a byte array, encode it in hexadecimal
    encode(sha256(random()::text::bytea), 'hex')
);

-- +goose Down
ALTER TABLE users DROP  COLUMN api_key;
