INSERT INTO role (title) VALUES ('Администратор');
INSERT INTO role (title) VALUES ('Архивариус');
INSERT INTO role (title) VALUES ('Гость');

INSERT INTO content_state (title) VALUES ('Архивировано');
INSERT INTO content_state (title) VALUES ('Удалено');

INSERT INTO content_action (title) VALUES ('Добавление');
INSERT INTO content_action (title) VALUES ('Обновление');
INSERT INTO content_action (title) VALUES ('Перемещение');
INSERT INTO content_action (title) VALUES ('Создание');
INSERT INTO content_action (title) VALUES ('Удаление');
INSERT INTO content_action (title) VALUES ('Удаление из короба');
INSERT INTO content_action (title) VALUES ('Переименование');

INSERT INTO account (name, username, email, password, role_id, active_status, branch)
VALUES ('admin', 'admin', 'admin@sadkomed.ru', '$2a$10$TDsjHTIQQbTwYGw6T1Ye4uGTrdiostY69m4nM5UeiLl9QcxcpEQSm', 1, true, 'sadkomed');


INSERT INTO box (code, content_state_id, unlimited_storage, description) VALUES ('S874210', 1, false,'test description');
INSERT INTO box (code, content_state_id, unlimited_storage, description) VALUES ('0005328', 1, false,'test description');
INSERT INTO box (code, content_state_id, unlimited_storage, description) VALUES ('S871561', 1, false,'test description');
INSERT INTO box (code, content_state_id, unlimited_storage, description) VALUES ('S414583', 1, false,'test description');


INSERT INTO record (archived_date, branch, pcode, last_treat, content_state_id, box_id) VALUES (now(), 'SADKO', 686755, now(), 1,1 );
INSERT INTO record (archived_date, branch, pcode, last_treat, content_state_id, box_id) VALUES (now(), 'SADKO', 584857, now(), 1,1 );
INSERT INTO record (archived_date, branch, pcode, last_treat, content_state_id, box_id) VALUES (now(), 'A2', 234554, now(), 1,2 );
INSERT INTO record (archived_date, branch, pcode, last_treat, content_state_id, box_id) VALUES (now(), 'SADKO', 198289, now(), 1,3 );
