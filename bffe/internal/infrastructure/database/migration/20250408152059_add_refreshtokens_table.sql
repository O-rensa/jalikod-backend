-- +goose Up
-- +goose StatementBegin
CREATE TABLE bffe.refresh_tokens (
    id SERIAL PRIMARY KEY,
    token VARCHAR(65) NOT NULL,
    associated_user_security_stamp UUID NOT NULL,
    user_id INTEGER NOT NULL,
    expires_on TIMESTAMP NOT NULL,
    is_revoked BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_refresh_tokens_token on bffe.refresh_tokens(token, is_revoked);

ALTER TABLE bffe.refresh_tokens
ADD CONSTRAINT fk_refresh_tokens_users
FOREIGN KEY (user_id) REFERENCES bffe.users(id)
ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE bffe.refresh_tokens
DROP CONSTRAINT fk_refresh_tokens_users;

DROP INDEX idx_refresh_tokens_token;

DROP TABLE bffe.refresh_tokens;
-- +goose StatementEnd
