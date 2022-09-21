-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION user_select(
    IN_ID INTEGER, IN_NAME VARCHAR, IN_USERNAME VARCHAR, IN_EMAIL VARCHAR, IN_ROLE_ID INT, IN_BRANCH VARCHAR
)
    RETURNS TABLE
            (
                id            INTEGER,
                name          VARCHAR,
                username      VARCHAR,
                email         VARCHAR,
                password      VARCHAR,
                branch        VARCHAR,
                role_id       INTEGER,
                role_title    VARCHAR,
                active_status BOOL,
                creation_date  BIGINT
            )

AS
$$
BEGIN
    RETURN QUERY SELECT a.id,
                        a.name,
                        a.username,
                        a.email,
                        a.password,
                        a.branch,
                        a.role_id,
                        r.title as role_title,
                        a.active_status,
                        extract(epoch from a.creation_date)::bigint AS creation_date
                 FROM account a
                          LEFT JOIN role r on r.id = a.role_id
                 WHERE (CASE IN_ID WHEN 0 THEN a.id > 0 ELSE a.id = IN_ID END)
                   AND (CASE IN_ROLE_ID WHEN 0 THEN a.role_id > 0 ELSE a.role_id = IN_ROLE_ID END)
                   AND (a.name ilike concat('%', IN_NAME, '%'))
                   AND (a.email ilike concat('%',IN_EMAIL,'%'))
                   AND (a.username ilike concat('%', IN_USERNAME, '%'))
                   AND (a.branch ilike concat('%', IN_BRANCH, '%'));


END
$$
    LANGUAGE 'plpgsql';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION user_select(IN_ID INTEGER, IN_NAME VARCHAR, IN_USERNAME VARCHAR, IN_EMAIL VARCHAR, IN_ROLE_ID INT, IN_BRANCH VARCHAR);
-- +goose StatementEnd


