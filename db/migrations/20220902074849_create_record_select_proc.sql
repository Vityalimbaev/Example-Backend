-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION record_select(
    IN_ID INTEGER, IN_PCODE BIGINT, IN_BRANCH VARCHAR, IN_BOX_ID INTEGER,
    IN_CONTENT_STATE_ID INT, START_ARCHIVED_DATE BIGINT, END_ARCHIVED_DATE BIGINT,
    START_CREATION_DATE BIGINT, END_CREATION_DATE BIGINT,
    START_LAST_TREAT BIGINT, END_LAST_TREAT BIGINT
)
    RETURNS TABLE
            (
                id               INTEGER,
                archived_date     BIGINT,
                branch           VARCHAR,
                creation_date    BIGINT,
                pcode            BIGINT,
                last_treat       BIGINT,
                content_state_id INTEGER,
                box_id           INTEGER
            )

AS
$$
BEGIN

    RETURN QUERY SELECT record.id,
                        extract(epoch from record.archived_date)::bigint AS archived_date,
                        record.branch,
                        extract(epoch from record.creation_date)::bigint AS creation_date,
                        record.pcode,
                        extract(epoch from record.last_treat)::bigint AS last_treat,
                        record.content_state_id,
                        record.box_id
                 FROM record
                 WHERE (CASE IN_ID WHEN 0 THEN record.id > 0 ELSE record.id = IN_ID END)
                   AND (CASE IN_PCODE WHEN 0 THEN record.pcode > 0 ELSE record.pcode = IN_PCODE END)
                   AND (CASE IN_CONTENT_STATE_ID WHEN 0 THEN record.content_state_id > 0 ELSE record.content_state_id = IN_CONTENT_STATE_ID END)
                   AND (CASE IN_BOX_ID WHEN 0 THEN record.box_id > 0 ELSE record.box_id = IN_CONTENT_STATE_ID END)
                   AND (record.branch ilike concat('%', IN_BRANCH, '%'))
                   AND (record.archived_date BETWEEN (to_timestamp(START_ARCHIVED_DATE) AT TIME ZONE 'UTC') AND  (to_timestamp(END_ARCHIVED_DATE) AT TIME ZONE 'UTC'))
                   AND (record.creation_date BETWEEN (to_timestamp(START_CREATION_DATE) AT TIME ZONE 'UTC') AND (to_timestamp(END_CREATION_DATE) AT TIME ZONE 'UTC'))
                   AND (record.last_treat BETWEEN (to_timestamp(START_LAST_TREAT) AT TIME ZONE 'UTC') AND (to_timestamp(END_LAST_TREAT) AT TIME ZONE 'UTC'));


END
$$
    LANGUAGE 'plpgsql';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS record_select;
-- +goose StatementEnd
