-- タスクを管理するテーブル
CREATE TABLE `tasks` (
  `id`      serial       COLLATE utf8mb4_bin NOT NULL COMMENT 'task id',
  `name`    varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'taskの名前',
  `status`  varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'taskのステータス',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='taskを管理するテーブル';