-- +goose Up
-- +goose StatementBegin
CREATE TABLE bffe.users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    middle_initial VARCHAR(10),
    surname VARCHAR(100) NOT NULL,
    name_extension VARCHAR(10),
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    security_stamp UUID,
    concurrency_stamp UUID NOT NULL,

    is_deleted BOOLEAN NOT NULL DEFAULT TRUE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    creation_time TIMESTAMP NOT NULL,
    creator_userid INT NOT NULL,
    last_modification_time TIMESTAMP,
    last_modifier_userid INT,
    deletion_time TIMESTAMP,
    deleter_userid INT
);

CREATE INDEX idx_users_concurrency_stamp ON bffe.users(concurrency_stamp);
CREATE INDEX idx_users_username ON bffe.users(username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_users_username;
DROP INDEX idx_users_concurrency_stamp;
DROP TABLE bffe.users;
-- +goose StatementEnd
