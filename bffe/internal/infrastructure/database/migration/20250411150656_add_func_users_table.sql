-- +goose Up
-- +goose StatementBegin
CREATE FUNCTION get_user_username(user_name VARCHAR)
RETURNS VARCHAR
LANGUAGE plpgsql
AS $$
DECLARE
    result_username VARCHAR;
BEGIN
    SELECT username INTO result_username 
    FROM bffe.users
    WHERE username = user_name;

    RETURN result_username;
END;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION get_user_username(user_name VARCHAR);
-- +goose StatementEnd
