# のみぱら 🍺

[![IMAGE ALT TEXT HERE](./doc/image/index/nomipara.png)]()

<br>

## 🎥 作品紹介動画

<!-- TODO -->
[Youtube 動画](https://youtu.be/JB3t8Gnhozg)

<br>

## ✨ 作品URL
https://jphacks-app.herokuapp.com

<br>

## 🖋 スライドURL
[Google スライド](https://docs.google.com/presentation/d/1ymClbAbEdxkUMlKGfkSrC91aTtMSw-LOrUh6XAYwOCw/edit?usp=sharing)

<br>

## 💪 製品概要

オンライン飲み会に最適化された新しいビデオ通話サービスです！

オンライン飲み会の課題を解決し、快適なオンライン飲み会体験を提供します。

<br>

### 🔥 背景 (製品開発のきっかけ、課題等)

皆さん、大人数での飲み会は好きですか？🍻

対面での飲み会は、いろんな人同士がさまざまな話題でグループを作って会話を楽しんでいました。

<br>

しかし、オンライン飲み会には、「**複数人が並行して会話することができない**」という大きな課題があります。

並行して会話ができないと以下のようなデメリットがあります。

<details>
<summary>大勢参加者がいたとしても、話す話題が一つになり、並行して会話を進めることができない</summary>

- ルームの参加者が強制的に発言者の話題に縛られてしまう

</details>

<details>
<summary>話すタイミングを見失う</summary>

- 同時に話し始めると気まずい
- 誰かが話し始めるのを待って沈黙の時間が続く
  - 悪魔の沈黙、誰か話してくれーー😢

</details>

<details>
<summary>ほかの参加者に気を使って好きな話題を話すことができない</summary>

- 多くの人がわからない推しの技術など
  - 自分はクリーンアーキテクチャの話が好きです👀 @mahiro72

</details>


<br>

対面時のように複数人が並行して会話することが難しいのです。

<br>

そこで私たちは、上記の悩みを解決し、

オンライン飲み会をさらに楽しくするサービス、「のみぱら🍺」を開発しました。

<br>

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

#### 使い方

大きく分けて３つのステップです。

1. 部屋を作成する

   [こちら](https://jphacks-app.herokuapp.com/room/prepare) からroom作成ページにアクセスし、好きな部屋番号を入力します (例:1111)

   入力したら作成ボタンを押して、部屋を作成します。

   部屋を作成したら、カメラをオンにしてほかの参加者を待ちましょう！

   ※下の写真はレッドブルを持っていますが、このサービスはビールが推奨です。

   <img src="./doc/image/index/step1.png" width="400px" />

2. 部屋番号を参加者に共有する

   部屋を作成する際に入力した番号を、参加者に共有します。

   参加者は[こちら](https://jphacks-app.herokuapp.com/room/prepare)にアクセスし、部屋番号を入力して参加を押します。

3. オンライン飲み会を楽しむ！🍻
   
   あとは楽しむだけです！！

   のみぱら！🍺

   <img src="./doc/image/index/party.png" width="400px" />

<br>

余談ですが

「のみぱら🍺」の、のみは飲み会の、のみですが、ぱらには２つの意味があります。

１つめは、parallel(並行)のぱらです。
並行して会話ができるサービスなので、パラレルのぱらを使いました。

２つめは、パラダイスのぱらです。
パラダイスは楽園などの意味があります。

ようは**思いっきり楽しんで飲み会をしましょう**ということです！！🍻

<br>

### 🥳 特長

#### 1. 一つのビデオ通話にいながら、別々の話題で会話することができる

##### フォーカス機能

フォーカスについては[こちら](https://github.com/jphacks/F_2205#%E3%83%95%E3%82%A9%E3%83%BC%E3%82%AB%E3%82%B9%E6%A9%9F%E8%83%BD)をご覧ください。

###### クリックでフォーカスする、解除する

クリックした人をフォーカスできます。

フォーカスするとオレンジの枠で接続している状態となり、

お互いの声が大きく聞こえ、ほかの人の声は小さく聞こえます。

<img src="./doc/image/index/focus2man.png" width="400px" />

<br>

さらにほかの人をクリックすると、フォーカスする人数を増やせます。

下記の状態では２人と接続している状態となっており、

２人の声が大きく聞こえます。

<img src="./doc/image/index/focus3man.png" width="400px" />

<br>

もう一度クリックしてフォーカス状態を解除できます。

<br>


左下の人がたくさんいるボタンを押すとすべてのフォーカスを解除できます。

誰もフォーカスしていない状態だと、参加者全員の声が同じ大きさで聞こえます。

<br>

この機能により、お互いに接続状態のユーザーは周りの声が気にならず

自由に会話を楽しむことができます。

<br>

#### 2. 視線操作で会話する相手を選択する

##### 視線でフォーカスする

右下の目のボタンを押すと、赤色の点が表示されます。

これはあなたの視線を追跡しています。

<img src="./doc/image/index/sisen_notmatch.png" width="400px" />


あなたの視線が、誰かを見ているかログをとり、

あなたがよくみているユーザーと接続状態にします。

<img src="./doc/image/index/sisen_match.png" width="400px" />

<br>


### ✅ 解決出来ること

- 1つのビデオ通話にいながら複数人が並行して会話することで、参加者みんなが会話を自由に楽しむことができる！
- オンライン飲み会をより楽しめるサービスを提供し、コロナ禍で減っていた人とのつながりの機会を提供する！

<br>

### ✨ 今後の展望

私たちの作った**のみぱら🍺**は既存の課題であった「**並行して会話することができないこと**」を解決しました。

そこで次の目標となるのは、さらに楽しいオンライン飲み会プラットフォームを提供することです。

具体的には、オンライン飲み会をさらに盛り上げるため、以下の機能を実装します。

<details>
<summary>チャット機能</summary>

- ユーザー同士のチャット用
- 定期的に話題を提供するbotの導入

</details>

<details>
<summary>スクショ機能</summary>

- スクショ時にエフェクトをつけて、飲み会が盛り上がっている様子を演出する

</details>

<details>
<summary>飲みすぎパラメータ(実装済み)</summary>

- お酒を飲む動作をした場合、ゲージが上昇していき、飲みすぎてないか常に確認できる
- 逆に飲んでない人もわかりますが、アルハラしないように気を付けましょう⚠️

</details>

<details>
<summary>一気飲み警告エフェクト(実装済み)</summary>

- お酒を飲む動作をした場合、そのユーザーに盛り上がっているエフェクトがつく
  - エフェクトが付きますが、お酒に強くなったわけではありません
  - お酒は楽しんで飲みましょう！🍻

</details>

<br>

### 😆 注力したこと（こだわり等）

- 並行して複数人が会話できる機能を追加し、参加者が自由に会話できる環境を提供できるようにした
- 視線を計測し、現実の飲み会のように自分がみているユーザーの声が大きく聞こえる
- weboskcet通信で、常に誰と誰がフォーカスしているかの情報を共有できるようにした
- DBに保存せずオンメモリでデータを管理することで、さらに高速に情報の共有をできるようにした
- サーバーサイドをクリーンアーキテクチャで実装したことで、DBを利用することになってもすぐに変更できるようにした
- デフォルトのクリックによる精度調整を行いやすいようなフォームとカルマンフィルタによるより正確な視点情報収集機能の作成により視覚トラッキングパッケージの精度向上を行った
- 自分の声は、一部のユーザー同士にしか大きく聞こえないようにできるため、各々が自由に話せる環境を作った


<br><br>

## 開発技術

### アーキテクチャ

<img src="https://user-images.githubusercontent.com/70263039/197279527-6e4a8def-e689-4a02-abe8-03ee3c84ea56.png" width="600px" />

<br>

### 活用した技術

#### API・データ

- [SkyWay](https://webrtc.ecl.ntt.com/)

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
    - Docker / Docker Compose
  - 本番環境
    - フロントエンド
      - Heroku
    - サーバーサイド
      - GCP
      - Cloud Run
      - Cloud Build
- CI/CD
  - GitHub Actions
  - heroku
  - Cloud Run / Cloud Build / Container Registry
- 開発環境整備
  - Make

<br>

#### デバイス

- PC (google chrome 推奨)

<br>

### 独自技術

#### ハッカソンで開発した独自機能・技術

##### フォーカス機能

フォーカス機能はユーザー同士を接続する機能です。

<br>

下記の図のように、

BさんからAさん、CさんからAさん、DさんからAさんにfocusのリクエストを送ると、

AさんとBさん、CさんとAさん、DさんとAさん、それぞれが接続している状態になります。

このとき、Aさんは全員と接続しているのでB,C,Dさんの声が通常通り聞こえますが、

B,C,DさんはAさんの声以外、小さく聞こえるようになります。

<img src="./doc/image/index/focus.png" width="400px" />

<br>

このように、誰かをフォーカスすることで、
その対象のユーザーと強制的にお互いにフォーカスしあっている状態とする機能がフォーカス機能です。

<br>

この機能を使うことで、
話したいユーザー以外の声を小さくし、好きな話題をそれぞれが並行して会話できる環境ができます。

<br>

つづいて具体例をみます。

この場合、AさんとBさんはお互いにフォーカスしあっており、
また、CさんとDさんもお互いにフォーカスしあっています。

このとき彼らはお互いの声以外は小さく聞こえており、
自分たちの好きな話題に集中して会話できます。

Eさんはだれともフォーカスしあってないので、
参加者全員の声が通常の大きさで聞こえます。


<img src="./doc/image/index/focus2.png" width="400px" />

<!-- TODO どこで実装されているか -->

<br>

#### 視線操作による会話相手決定機能

こちらは、自分が話したいと思った相手をみることでその人に上記のフォーカス機能が適用される機能です。 

今回使用したWebGazerは一般的な Web カメラを使用して、ページ上の Web 訪問者の視線位置をリアルタイムで推測するアイトラッキングライブラリです。 

今回自分が話たいと思った相手を認識させるために以下のような手順を踏む処理を構築しました。

1. webgazerを起動し,ユーザーの視線みている場所を現在使っているブラウザからみた相対値のx,y座標として数ミリ秒ごとに取得する

2. 取得したx,v座標から現在みているHTMLのelement要素を取得する 

3. 2で取得したelement要素が10回連続ユーザー画面を表すもの稼働かを判定する。(サーバー負荷対策)

4. 3でtrueになった場合上記のフォーカス機能を走らせる 

##### webgazer認識精度調整機能

またwebgazerは初期状態では精度に課題を抱えていました。 

そのためデフォルトで備わっているマウスクリックによるフィードバック修正に加えて、カルマンフィルタを使った視点認識の調整機能、

収集した視点データから現在のwebgazerの認識精度の計算表示機能を実装しました。 

またこの収集データをlocalstorageに保存し二回目以降のみぱらを立ち上げた際に参照することで、

認識能力を維持する機能もあわせて実装しました。 

<br>



#### 製品に取り入れた研究内容（データ・ソフトウェアなど）（※アカデミック部門の場合のみ提出必須）

- カルマンフィルタ
  - 参考 : https://inzkyk.xyz/kalman_filter/designing_kalman_filters/ 

[WebGazer: Scalable Webcam Eye Tracking Using User Interactions](https://cs.brown.edu/people/apapouts/papers/ijcai2016webgazer.pdf)
- [MediaPipe Facemesh](https://google.github.io/mediapipe/solutions/face_mesh.html)
- [WebGazer (GitHub)](https://github.com/brownhci/WebGazer)

<br>


### プロダクト開発

#### MVP,Usecaseの検討

私たちのチームはMiroで事前にMVPやUsecaseを検討したことで、

必要な機能をしっかりと選定したうえで開発を進めることができました。

<img src="./doc/image/index/usecase.png" width="600px" />

<img src="./doc/image/index/mvp.png" width="600px" />

<img src="./doc/image/index/mvp2.png" width="600px" />

<br>

#### デザイン

また、Figmaを用いてデザインを行い、

チーム内でのプロダクトのデザインの共通認識を行ってスムーズに開発を進めました。

<img src="./doc/image/index/figma.png" width="600px" />

<br>

#### タスク管理

私たちのチームではタスクをISSUEで管理していました。

ISSUEのラベルを使うことで、タスクの優先度や分野について、わかりやすいように工夫しました。

<img src="./doc/image/index/issue.png" width="600px" />

<br>

#### 作業・進捗管理

Discordのボットを利用することで、

誰が今どの作業をしているのかを全体で共有しました。

<img src="./doc/image/index/discord.png" width="400px" />

<br>

### Miro

この [miro.md](./doc/miro.md) にmiroに書いた設計などをまとめてあります。

### おしまい

![export_image_5029](https://user-images.githubusercontent.com/29572313/201506655-613cec99-0ba0-4138-ac1f-c08d4581645a.png)