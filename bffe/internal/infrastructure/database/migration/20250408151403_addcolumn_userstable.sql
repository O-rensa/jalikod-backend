-- +goose Up
-- +goose StatementBegin
ALTER TABLE bffe.users
ADD COLUMN password_reset_expiry TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE bffe.users
DROP COLUMN password_reset_expiry;
-- +goose StatementEnd
