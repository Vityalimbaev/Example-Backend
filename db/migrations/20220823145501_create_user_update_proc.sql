-- +goose Up
-- +goose StatementBegin
-- CREATE OR REPLACE FUNCTION user_update(
--     in_id int, in_name varchar, in_username varchar,  in_email varchar, in_password varchar, branch varchar, role_id int, in_active_status bool

CREATE OR REPLACE PROCEDURE user_update (
    in_id int,
    in_name varchar,
    in_username varchar,
    in_email varchar,
    in_password varchar,
    in_branch varchar,
    in_role_id int
)
    language plpgsql
as
$$
begin

    UPDATE account
    set name     = COALESCE(NULLIF(in_name, ''), name),
        username = COALESCE(NULLIF(in_username, ''), username),
        email    = COALESCE(NULLIF(in_email, ''), email),
        password = COALESCE(NULLIF(in_password, ''), password),
        branch   = COALESCE(NULLIF(in_branch, ''), branch),
        role_id  = COALESCE(NULLIF(in_role_id, 0), role_id)

    where id = in_id;

    commit;
end ;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS user_update(in_id int, in_name varchar, in_username varchar, in_email varchar, in_password varchar, in_branch varchar, in_role_id int);
-- +goose StatementEnd
