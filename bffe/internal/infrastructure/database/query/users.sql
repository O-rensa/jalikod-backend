-- name: RegisterUserAndGetId :one
INSERT INTO bffe.users (
    first_name,
    middle_initial,
    surname,
    name_extension,
    username,
    password,
    security_stamp,
    concurrency_stamp,
    is_deleted,
    is_active,
    creation_time,
    last_modification_time,
    last_modifier_userid,
    deletion_time,
    deleter_userid
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
$11, $12, $13, $14, $15)
RETURNING id;

-- name: GetUserUsername :one
SELECT get_user_username($1) AS username;