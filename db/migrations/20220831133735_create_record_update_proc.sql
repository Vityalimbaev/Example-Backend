-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE PROCEDURE record_update (
    in_id int,
    in_archived_date timestamptz,
    in_branch varchar,
    in_pcode bigint,
    in_last_treat timestamptz
)
    language plpgsql
as
$$
begin

    UPDATE record
    set archived_date     = COALESCE(NULLIF(in_archived_date, to_timestamp(0)), archived_date),
        branch    = COALESCE(NULLIF(in_branch, ''), branch),
        pcode = COALESCE(NULLIF(in_pcode, 0), pcode),
        last_treat  = COALESCE(NULLIF(in_last_treat, to_timestamp(0)), last_treat)
    where id = in_id;
    commit;
end
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS record_update(in_id int, in_archived_date timestamp, in_branch varchar, in_pcode bigint, in_last_treat timestamp);
-- +goose StatementEnd
