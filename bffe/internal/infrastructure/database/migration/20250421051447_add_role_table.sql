-- +goose Up
-- +goose StatementBegin
CREATE TABLE bffe.roles (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(100) NOT NULL,

    is_deleted BOOLEAN NOT NULL DEFAULT TRUE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    creation_time TIMESTAMP NOT NULL,
    creator_userid INT NOT NULL,
    last_modification_time TIMESTAMP,
    last_modifier_userid INT,
    deletion_time TIMESTAMP,
    deleter_userid INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE bffe.roles;
-- +goose StatementEnd
