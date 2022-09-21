-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_session (
    id SERIAL NOT NULL,
    refresh_token VARCHAR(255) NOT NULL UNIQUE,
    date_expire bigint,
    account_id INT NOT NULL UNIQUE,

    PRIMARY KEY (id),
    CONSTRAINT user_session_account_account_id_foreign FOREIGN KEY (account_id) REFERENCES account (id)

) WITHOUT OIDS;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_session;
-- +goose StatementEnd
