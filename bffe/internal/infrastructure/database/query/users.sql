-- name: RegisterUserAndGetId :one
INSERT INTO bffe.users (
    first_name,
    middle_initial,
    surname,
    name_extension,
    email,
    username,
    password,
    password_reset_code,
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
$11, $12, $13, $14, $15, $16, $17)
RETURNING id;