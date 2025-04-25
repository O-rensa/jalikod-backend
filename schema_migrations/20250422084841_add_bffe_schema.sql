-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA bffe;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA bffe;
-- +goose StatementEnd
