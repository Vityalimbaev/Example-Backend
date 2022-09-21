-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE PROCEDURE box_update (
    in_id INT,
    in_code VARCHAR,
    in_content_state_id VARCHAR,
    in_unlimited_storage BOOL,
    in_description VARCHAR
)
    LANGUAGE plpgsql
AS
$$
BEGIN

    UPDATE box
    SET code     = COALESCE(NULLIF(in_code, ''), code),
        content_state_id = COALESCE(NULLIF(in_content_state_id, 0), content_state_id),
        unlimited_storage = in_unlimited_storage,
        description  = COALESCE(NULLIF(in_description,'' ), description)
    WHERE id = in_id;
    COMMIT;
END
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS record_update(in_id INT, in_archived_date TIMESTAMP, in_branch VARCHAR, in_pcode BIGINT, in_last_treat TIMESTAMP);
-- +goose StatementEnd