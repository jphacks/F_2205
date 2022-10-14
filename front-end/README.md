# 起動方法

##### コンテナをビルド・起動

```
docker-compose up -d
```

##### 依存関係をインストール

```
docker-compose exec nuxt yarn install
```

##### 開発用サーバを起動

```
docker-compose exec nuxt yarn dev
```

## vuetifyjs ドキュメント

[https://vuetifyjs.com/ja/](https://vuetifyjs.com/ja/)

## デプロイ

master ブランチに push することで heroku に自動デプロイされます。  
/front-end/nuxt-app 配下のファイルが heroku にデプロイされます。(back-end のプログラムは含みません)

##### 本番環境 URL

[https://jphacks-app.herokuapp.com/](https://jphacks-app.herokuapp.com/)
