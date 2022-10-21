# のみぱら 🍺

[![IMAGE ALT TEXT HERE](https://user-images.githubusercontent.com/70263039/197263117-0f6c1a18-c6f8-44b0-989e-1ae611e7c907.png)]()

<br>

## 🎥 作品紹介動画
https://www.youtube.com/hogehoge

<br>

## ✨ 作品URL
https://jphacks-app.herokuapp.com/room/prepare

<br>

## 💪 製品概要

オンライン飲み会に最適化された新しいビデオ通話サービスです！

オンライン飲み会の課題を解決し、快適なオンライン飲み会体験を提供します。

<br>

### 🔥 背景 (製品開発のきっかけ、課題等)

オンライン飲み会は、対面の飲み会より盛り上がりにくいと感じたことはありませんか？


私たちは、対面での飲み会と比較し、オンライン飲み会の抱える大きな課題として

「**並行して会話することができないこと**」をあげます。

並行して会話ができないと以下のようなデメリットがあります。

- 大勢参加者がいたとしても、話す話題が一つになり、並行して会話を進めることができない
  - ルームの参加者が強制的に発言者の話題に縛られてしまう
- 話すタイミングを見失う
  - 同時に話し始めると気まずい
  - 誰かが話し始めるのを待って沈黙の時間が続く
    - 悪魔の沈黙、誰か話してくれーー😢
- ほかの参加者に気を使って好きな話題を話すことができない
  - 多くの人がわからない推しの技術など
    - 自分はクリーンアーキテクチャの話が好きです👀 @mahiro72

上記の悩みを解決し、オンライン飲み会をさらに楽しくするためアプリを開発しました。

<br>

<!-- TODO 使うか検討 -->
<!-- #### 【 開発のきっかけ✨ 】 飲み会は「人と人との繋がり」を作ることができる

コロナ渦になり、人との交流が難しい社会情勢になってしまいました。

特に大人数での飲み会は感染リスクが高く自粛されることが多いでしょう。 

そして、オンラインでの飲み会は人数が増えるほど盛り上がるのが難しいことが多いです。

ですが、飲み会は「人と人との繋がり」を作るのにとても有用です。

そこで、飲み会をオンラインでも対面と変わらない感覚で楽しめるサービスが必要だと考え、開発しました。

<br> -->

### 📝 製品説明（具体的な製品の説明）

既存のビデオ通話機能に加え、一つのビデオ通話にいながら、並行して別々の話題で会話することができるサービスです！

また、視線トラッキングを活用することでより対面でのオンライン飲み会体験に近づけています。

<br>

余談ですが

のみぱらの、のみは飲み会の、のみですが、ぱらには２つの意味があります。

１つめは、parallel(並行)のぱらです。
平行して会話ができるサービスなので、パラレルのぱらを使いました。

２つめは、パラダイスのぱらです。
パラダイスは楽園などの意味があります。

ようは**思いっきり楽しんで飲み会をしましょう**ということです！！🍻

<br>
### 🥳 特長

#### 1. 一つのビデオ通話にいながら、別々の話題で会話することができる

<!-- TODO 写真欲しいな -->

<br>

#### 2. 視線操作で会話する相手を選択する

<!-- TODO 写真欲しいな -->

<br>

#### 3. 特長 3

<!-- TODO 写真欲しいな -->

<br>

### ✅ 解決出来ること

<!-- TODO 書き方がいまいちなので直したい -->

- 1つのビデオ通話にいながら並行して会話し、参加者みんなが盛り上がることができる！
- 視線を計測し、現実の飲み会のように自分がみているユーザーの声が大きく聞こえる！
- 自分の声は、一部のユーザー同士にしか大きく聞こえないようにできるため、好きな話題を好きなだけ話せる！
- オンライン飲み会をより楽しめるサービスを提供し、人とのつながりの機会を提供する！

<br>

### ✨ 今後の展望

私たちの作った**のみぱら🍺**は既存の課題であった「**並行して会話することができないこと**」を解決しました。

そこで次の目標となるのは、さらに楽しいオンライン飲み会プラットフォームを提供することです。

具体的には、オンライン飲み会をさらに盛り上げるため、以下の機能を実装します。

- チャット機能
  - ユーザー同士のチャット用
  - 定期的に話題を提供するbotの導入
- スクショ機能
  - スクショ時にエフェクトをつけて、飲み会が盛り上がっている様子を演出する
- 飲みすぎパラメータ
  - お酒を飲む動作をした場合、ゲージが上昇していき、飲みすぎてないか常に確認できる
  - 逆に飲んでない人もわかりますが、アルハラしないように気を付けましょう⚠️
- 一気飲みエフェクト
  - お酒を飲む動作をした場合、そのユーザーに盛り上がっているエフェクトがつく
    - エフェクトが付きますが、お酒に強くなったわけではありません
    - お酒は楽しんで飲みましょう！🍻

<br>

### 😆 注力したこと（こだわり等）

<!-- TODO もっとしっかりかく -->
- 視線まわり
- ビデオ通話
- websocketでよりリアルタイムな通信

<br><br>

## 開発技術

### アーキテクチャ

<img src="https://user-images.githubusercontent.com/70263039/197279031-46a86707-334e-4378-81b0-dc9f1986902d.png" width="600px" />


### 活用した技術

#### API・データ

<!-- TODO 説明入れる -->
- SkyWay

<br>

#### フレームワーク・ライブラリ・モジュール

- フロントエンド
  - Nuxt.js
    - HTML,SCSS
    - JavaScript
  - Prettier
  - vuetifyjs
- サーバーサイド
  - Golang
  - [gorilla/websocket](https://github.com/gorilla/websocket)
  - [gin-gonic/gin](https://github.com/gin-gonic/gin)
- インフラ
  - 開発環境
    - Docker
    - Docker Compose
  - 本番環境
    - フロントエンド
      - Heroku
    - サーバーサイド
      - GCP
      - Cloud Run
      - Cloud Build
  - GitHub Actions
- CI/CD
  - GitHub Actions
  - heroku
  - Cloud Run
- 開発環境整備
  - Make

<br>

#### デバイス

- PC (google chrome 推奨)

<br>

### 独自技術

- ...
- ...

<br>

#### ハッカソンで開発した独自機能・技術

##### フロントエンド

- 独自で開発したものの内容をこちらに記載してください
- 特に力を入れた部分をファイルリンク、または commit_id を記載してください。

<br>

##### サーバーサイド

- hoge

<br>

#### 製品に取り入れた研究内容（データ・ソフトウェアなど）（※アカデミック部門の場合のみ提出必須）

-
-

<br>
