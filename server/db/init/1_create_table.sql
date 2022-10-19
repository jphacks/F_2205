-- DB切り替え
\c "hoge"

--テーブルを作成

CREATE TABLE "rooms" (
  "id"         SERIAL NOT NULL PRIMARY KEY,
  "name"       VARCHAR(255) NOT NULL
);
