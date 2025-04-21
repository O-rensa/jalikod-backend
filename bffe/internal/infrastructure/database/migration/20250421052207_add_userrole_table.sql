-- +goose Up
-- +goose StatementBegin
CREATE TABLE bffe.user_roles (
    id SERIAL PRIMARY KEY,
    userid int NOT NULL,
    roleid int NOT NULL
);

ALTER TABLE bffe.user_roles
ADD CONSTRAINT fk_userRoles_users
FOREIGN KEY (userid)
REFERENCES bffe.users(id);

ALTER TABLE bffe.user_roles
ADD CONSTRAINT fk_userRoles_roles
FOREIGN KEY (roleid)
REFERENCES bffe.roles(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE bffe.user_roles
DROP CONSTRAINT fk_userRoles_roles;

ALTER TABLE bffe.user_roles
DROP CONSTRAINT fk_userRoles_users;

DROP TABLE bffe.user_roles;
-- +goose StatementEnd
