# Server

## ディレクトリ構成(/server)
```
|--.env.local  # localの環境変数
|--.gitignore
|--db          # database
|--doc         # documents
|--docker      # docker composeファイル
|--html        # serverテスト用html
|--Makefile
|--README.md
|--src         # serverのソースコード
```

## ディレクトリ構成(/server/src)

```
|--config          # 接続などの設定回り
|--domain
|  |--entity       # ドメインのオブジェクトを管理します
|  |--repository   # infrastructure/repositoryのインターフェースを管理します
|  |--service
|--infrastructure
|  |--database     # データベースと接続します
|  |--persistance  # データの永続化を責務とします
|--main.go
|--presentation
|  |--handler      # ハンドラーを管理します
|  |--router       # ルーターを管理します
|  |--ws           # websocket通信を管理します
|--usecase         # ユースケースを管理します
|--utils           # 共通で使う関数をまとめてあります
```

## API仕様書

APIについてまとめてあります

API仕様書は[こちら](./doc/api.md)から

