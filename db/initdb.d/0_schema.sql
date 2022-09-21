

CREATE TABLE IF NOT EXISTS role
(
    id    SERIAL       NOT NULL,
    title VARCHAR(255) NOT NULL UNIQUE,
    PRIMARY KEY (id)
) WITHOUT OIDS ;

CREATE TABLE IF NOT EXISTS content_state
(
    id    SERIAL       NOT NULL,
    title VARCHAR(255) NOT NULL UNIQUE,
    PRIMARY KEY (id)
) WITHOUT OIDS;

CREATE TABLE IF NOT EXISTS content_action
(
    id    SERIAL       NOT NULL,
    title VARCHAR(255) NOT NULL UNIQUE,
    PRIMARY KEY (id)
) WITHOUT OIDS;

CREATE TABLE IF NOT EXISTS account
(
    id            SERIAL       NOT NULL,
    name          VARCHAR(255) NOT NULL,
    email         VARCHAR(255) NOT NULL CONSTRAINT account_email_unique UNIQUE,
    branch        VARCHAR(255),
    role_id       INT          NOT NULL,
    active_status BOOLEAN DEFAULT FALSE,
    username      VARCHAR(255) NOT NULL  CONSTRAINT account_username_unique UNIQUE,
    password      VARCHAR(255),
    creation_date TIMESTAMP DEFAULT now(),

    PRIMARY KEY (id),
    CONSTRAINT account_role_role_id_foreign FOREIGN KEY (role_id) REFERENCES role (id)
) WITHOUT OIDS;

CREATE TABLE IF NOT EXISTS user_session
(
    id                   SERIAL       NOT NULL,
    refresh_token                 VARCHAR(255) NULL,
    date_expire        bigint    NOT NULL ,
    account_id     INT          NOT NULL,

    PRIMARY KEY (id),
    CONSTRAINT user_session_account_account_id FOREIGN KEY (account_id) REFERENCES account (id)
) WITHOUT OIDS;


CREATE TABLE IF NOT EXISTS box
(
    id                   SERIAL       NOT NULL,
    code                 VARCHAR(255) NULL  CONSTRAINT box_code_unique UNIQUE,
    creation_date        TIMESTAMP    NULL DEFAULT now(),
    content_state_id     INT          NOT NULL,
    unlimited_storage BOOLEAN      NOT NULL,
    description          VARCHAR(255) NULL,

    PRIMARY KEY (id),
    CONSTRAINT box_content_state_content_state_id_foreign FOREIGN KEY (content_state_id) REFERENCES content_state (id)
) WITHOUT OIDS;


CREATE TABLE IF NOT EXISTS record
(
    id               SERIAL         NOT NULL,
    archived_date    TIMESTAMP      NOT NULL,
    branch           VARCHAR(255)   NOT NULL,
    creation_date    TIMESTAMP      NOT NULL DEFAULT now(),
    pcode            BIGINT         NOT NULL,
    last_treat       timestamp,
    content_state_id INT,
    box_id           INT NOT NULL,

    PRIMARY KEY (id),
    CONSTRAINT record_content_state_content_state_id_foreign FOREIGN KEY (content_state_id) REFERENCES content_state (id),
    CONSTRAINT record_box_box_id_foreign FOREIGN KEY (box_id) REFERENCES box (id)
) WITHOUT OIDS;

CREATE TABLE IF NOT EXISTS history
(
    id                SERIAL NOT NULL,
    content_action_id INT    NOT NULL,
    box_id            BIGINT,
    record_id         BIGINT,
    datetime          TIMESTAMP DEFAULT now(),
    description       VARCHAR,
    account_id        BIGINT,

    PRIMARY KEY (id),
    CONSTRAINT history_content_action_content_action_id_foreign FOREIGN KEY (content_action_id) REFERENCES content_action (id)

) WITHOUT OIDS;
