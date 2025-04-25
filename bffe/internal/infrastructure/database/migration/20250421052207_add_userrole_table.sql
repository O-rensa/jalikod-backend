-- +goose Up
-- +goose StatementBegin
CREATE TABLE bffe.user_roles (
    id SERIAL PRIMARY KEY,
    userid int NOT NULL,
    roleid int NOT NULL
);

ALTER TABLE bffe.user_roles
ADD CONSTRAINT fk_user_roles_users
FOREIGN KEY (userid)
REFERENCES bffe.users(id)
ON DELETE CASCADE;

ALTER TABLE bffe.user_roles
ADD CONSTRAINT fk_user_roles_roles
FOREIGN KEY (roleid)
REFERENCES bffe.roles(id)
ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE bffe.user_roles
DROP CONSTRAINT fk_user_roles_roles;

ALTER TABLE bffe.user_roles
DROP CONSTRAINT fk_user_roles_users;

DROP TABLE bffe.user_roles;
-- +goose StatementEnd
