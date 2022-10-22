<template>
  <section>
    <section id="video-wrap" class="video">
      <video id="my-video" class="video-individual" autoplay muted playsinline></video>
    </section>

    <!-- ビデオステータスバー -->
    <div class="status-bar">
      <VideoState
        :leavingFn="this.roomLeaving"
        :gazeEstimatingFn="this.swtichEstimateGaze"
        :isEnableGazeEstimating="this.isEnableGazeEstimating"
        :focusThisVideoAllLiftFn="this.focusThisVideoAllLift"
        :videoMuteFn="this.videoMute"
        :audioMuteFn="this.audioMute"
        :myVideoStatus="this.myVideoStatus"
        :myAudioStatus="this.myAudioStatus"
      />
    </div>
    <!-- ビデオステータスバー -->

    <!-- モーダルウィンドウ -->
    <section class="modal-window">
      <div class="modal-window-back"></div>
      <div class="modal-window-front">
        <h3>部屋に接続する</h3>
        <div class="modal-window-front_btn-wrap">
          <Btn text="接続" color="orange" :clickedfn="this.roomConnection" />
        </div>
      </div>
    </section>
    <!-- モーダルウィンドウ -->
  </section>
</template>

<script>
import Peer from 'skyway-js';
import axios from 'axios';

import VideoState from '~/components/presentational/organisms/videoState';
import Btn from '~/components/presentational/atoms/btn';

import webgazer from 'webgazer';

export default {
  components: {
    VideoState,
    Btn
  },

  data() {
    return {
      APIKey: '5152bad7-4798-40b1-986a-a7e8f164a8a3',
      unfocusedVolume: 0.15,
      localStream: null,
      peer: null,
      websocketConn: null,
      roomMemberNum: 1,
      isVisibleSwitchButton: false,
      isEnableGazeEstimating: false,
      isFirstGazeEstimating: true,
      elementUnderGazeCount: 0,
      myVideoStatus: true,
      myAudioStatus: true
    };
  },

  methods: {
    setWebsocketEventListener: function (websocketConn) {
      websocketConn.onopen = function (e) {
        console.log('websocket connection');
      };
      websocketConn.onmessage = function (evt) {
        //フォーカス処理

        const unfocusedVolume = 0.15;

        //-------------------------------------------------------------------------//
        //フォーカス全解除&音量下げ
        const focusThisVideoAllLift = () => {
          const videos = document.querySelectorAll('.video-individual');

          for (let video of videos) {
            //音量設定
            video.volume = unfocusedVolume;
            video.classList.remove('video-individual-focus');
          }
        };
        //-------------------------------------------------------------------------//
        //-------------------------------------------------------------------------//
        //引数のメンバーをフォーカスしてそれ以外を除外
        const doFocus = (focusMemberData) => {
          if (focusMemberData.length == 0) {
            const videos = document.querySelectorAll('.video-individual');

            for (let video of videos) {
              //音量設定
              video.volume = 1;
            }
            return;
          }

          document.querySelector('#my-video').classList.add('video-individual-focus');

          for (let tgMemberData of focusMemberData) {
            let videoDom = document.getElementById(tgMemberData.name);
            videoDom.classList.add('video-individual-focus');
            videoDom.volume = 1;
          }
        };
        //-------------------------------------------------------------------------//

        const data = JSON.parse(evt.data);
        const myPeerId = document.querySelector('#my-video').getAttribute('name');

        if (data.type != 'SET_FOCUS' && data.type != 'DEL_ALL_FOCUS' && data.type != 'DEL_FOCUS') return;

        focusThisVideoAllLift();

        const memberDatas = data.focus.members;

        for (let memberData of memberDatas) {
          if (memberData.name == myPeerId) {
            //引数のメンバーをフォーカスしてそれ以外を除外
            doFocus(memberData.Connects);
          }
        }
      };
      websocketConn.onclose = function (evt) {
        console.log('websocket connection closed');
      };
      websocketConn.onerror = function (evt) {
        console.log('websocket error: ' + evt);
      };
    },
    setSkywayEventListener: function (mediaConnection) {
      mediaConnection.on('stream', (stream) => {
        //video要素にカメラ映像をセットして再生
        this.addVideo(stream);
        const videoElm = document.getElementById(stream.id);
        videoElm.srcObject = stream;
      });

      mediaConnection.on('open', () => {
        document.querySelector('#my-video').setAttribute('name', this.peer.id);
      });

      //ルームメンバーが退出したときに発火
      mediaConnection.on('peerLeave', (peerId) => {
        this.removeVideo(peerId);
      });

      //何らかのエラーが発生した場合に発火
      this.peer.on('error', (err) => {
        alert(err.message);
      });
    },

    roomConnection: function () {
      //ルーム作成 or ルーム参加
      console.log('my peer id: ' + this.peer.id);
      const roomName = this.$route.params.id;
      const mediaConnection = this.peer.joinRoom(roomName, { mode: 'sfu', stream: this.localStream });
      this.setSkywayEventListener(mediaConnection);
      document.querySelector('body').classList.remove('modal-open');
      this.isVisibleSwitchButton = true;

      //人数制限チェック
      setTimeout(this.roomMemberNumCheck, 5000);
      setInterval(this.roomMemberNumCheck, 60000);

      //ルーム接続時間制限(5分)
      setTimeout(this.roomLeaving, 300000);

      const data = {
        type: 'NEW_MEMBER',
        info: {
          from: `${this.peer.id}`
        }
      };

      this.websocketConn.send(JSON.stringify(data));
    },

    roomLeaving: async function () {
      //ルーム退出
      console.log('room leaving start');

      this.peer.destroy();

      //websocket そのユーザーの持っている接続状態を解除する
      if (this.roomMemberNum != 1) {
        const data = {
          type: 'DEL_ALL_FOCUS',
          info: {
            from: this.peer
          }
        };

        this.websocketConn.send(data);
      } else {
        const response = await axios.delete(
          'https://f-2205-server-chhumpv4gq-de.a.run.app/ws/' + this.$route.params.id
        );
        console.log(response);
      }

      alert('退出しました');

      //websocketの接続を切断(正常終了)
      this.websocketConn.close(1000, 'normal amputation websocket');

      //リダイレクト
      this.$router.push('/room/prepare');

      this.endEstimateGaze();
      this.isVisibleSwitchButton = false;
    },

    addVideo: function (stream) {
      const videoDom = document.createElement('video');
      videoDom.setAttribute('id', stream.peerId);
      videoDom.classList.add('video-individual');
      videoDom.srcObject = stream;
      videoDom.play();
      document.getElementById('video-wrap').append(videoDom);

      //ルームメンバー人数追加
      this.roomMemberNum++;

      //ビデオのリサイズ
      this.videoResize();
    },
    removeVideo: function (peerId) {
      console.log(peerId);
      const videoDom = document.getElementById(peerId);
      videoDom.remove();

      //ルームメンバー人数減少
      this.roomMemberNum--;

      //ビデオのリサイズ
      this.videoResize();
    },
    videoResize: function () {
      //ビデオのリサイズ
      const videos = document.querySelectorAll('.video-individual');

      for (let video of videos) {
        switch (this.roomMemberNum) {
          case 1:
            video.style.width = '100%';
            video.style.height = '100%';
            break;
          case 2:
            video.style.width = '38%';
            video.style.height = 'auto';
            break;
          default:
            video.style.width = '38%';
            video.style.height = 'auto';
            break;
        }
      }
    },

    beginEstimateGaze: function () {
      //視線からフォーカス
      webgazer
        .showVideo(false)
        .showPredictionPoints(true)
        .setGazeListener((gaze, clock) => {
          if (gaze == null) {
            return;
          }

          const x = gaze.x;
          const y = gaze.y;

          const elementUnderGaze = document.elementFromPoint(x, y);

          if (elementUnderGaze === null) return;

          if (elementUnderGaze.tagName == 'VIDEO') {
            this.elementUnderGazeCount++;
            // TODO: 試験的にカウントを10以上に設定, 後ほど適切な値・実装方法に変える
            if (this.elementUnderGazeCount > 10) {
              console.log('elementUnderGazeCount is 10 count');
              this.focusThisVideoLineOfSight(elementUnderGaze.id);
            }
          } else {
            this.elementUnderGazeCount = 0;
            //フォーカス全外し(この関数を呼ぶことでサーバー側にリクエスト飛ぶ)
            //this.focusThisVideoAllLift();
          }
        })
        .begin();
    },

    endEstimateGaze: function () {
      console.log('endEstimateGaze');
      webgazer.clearGazeListener().end();

      // webgazerをendしても視線予測のポインターが消えないため、直接Elementを削除
      const gazeDotEl = document.getElementById('webgazerGazeDot');
      gazeDotEl.remove();

      this.isFirstGazeEstimating = false;
    },

    pauseEstimateGaze: function () {
      console.log('pauseEstimateGaze');
      webgazer.showPredictionPoints(false).pause();
    },

    resumeEstimateGaze: function () {
      console.log('resumeEstimateGaze');
      webgazer.showPredictionPoints(true).resume();
    },

    swtichEstimateGaze: function () {
      this.isEnableGazeEstimating = !this.isEnableGazeEstimating;
    },

    focusThisVideo: function (id) {
      //ビデオをフォーカスする(自分のビデオ以外)
      if (id == 'my-video') return;

      const className = document.getElementById(id).className;
      if (className == 'video-individual') {
        //websocket ユーザー同士を接続状態にする
        const data = {
          type: 'SET_FOCUS',
          info: {
            from: `${this.peer.id}`,
            to: `${id}`
          }
        };
        this.websocketConn.send(JSON.stringify(data));
      } else {
        //websocket ユーザー同士を接続状態にする
        const data = {
          type: 'DEL_FOCUS',
          info: {
            from: `${this.peer.id}`,
            to: `${id}`
          }
        };
        this.websocketConn.send(JSON.stringify(data));
      }
    },
    focusThisVideoLineOfSight: function (id) {
      //ビデオをフォーカスする(自分のビデオ以外)(視線で)
      if (id == 'my-video') return;

      //websocket ユーザー同士を接続状態にする
      const data = {
        type: 'SET_FOCUS',
        info: {
          from: `${this.peer.id}`,
          to: `${id}`
        }
      };
      this.websocketConn.send(JSON.stringify(data));
    },
    focusThisVideoAllLift: function () {
      //フォーカスを全解除

      const data = {
        type: 'DEL_ALL_FOCUS',
        info: {
          from: `${this.peer.id}`
        }
      };
      this.websocketConn.send(JSON.stringify(data));

      const videos = document.querySelectorAll('.video-individual');

      for (let video of videos) {
        //音量設定
        video.volume = 1;
        video.classList.remove('video-individual-focus');
      }
    },
    roomMemberNumCheck: function () {
      //部屋の最大人数のチェック
      const roomMaxmemberNum = 18;

      console.log('check room member num');
      if (this.roomMemberNum > roomMaxmemberNum) {
        console.log('forced exit');
        this.roomLeaving();
      }
      console.log('ok');
    },
    videoMute: function () {
      //画面をミュート
      if (this.myVideoStatus) {
        this.localStream.getVideoTracks()[0].enabled = false;
      } else {
        this.localStream.getVideoTracks()[0].enabled = true;
      }
      this.myVideoStatus = !this.myVideoStatus;
    },
    audioMute: function () {
      //音声をミュート
      if (this.myAudioStatus) {
        this.localStream.getAudioTracks()[0].enabled = false;
      } else {
        this.localStream.getAudioTracks()[0].enabled = true;
      }
      this.myAudioStatus = !this.myAudioStatus;
    }
  },

  mounted: async function () {
    if (!this.$route.params.id) {
      alert('部屋番号が入力されていません');
      this.$router.push('/room/prepare');
    }

    //モーダルウィンドウを表示
    document.querySelector('body').classList.add('modal-open');

    //WebSocketで接続
    this.websocketConn = new WebSocket('wss://f-2205-server-chhumpv4gq-de.a.run.app/ws/' + this.$route.params.id);
    this.setWebsocketEventListener(this.websocketConn);

    //ビデオ設定(解像度を落とす)
    let constraints = {
      video: {},
      audio: true
    };
    constraints.video.width = {
      min: 320,
      max: 320
    };
    constraints.video.height = {
      min: 240,
      max: 240
    };

    try {
      // カメラ映像取得
      const stream = await navigator.mediaDevices.getUserMedia(constraints);
      document.getElementById('my-video').srcObject = stream;
      this.localStream = stream;
    } catch (error) {
      console.log(error);
    }

    //Peer作成
    this.peer = new Peer({
      key: this.APIKey,
      debug: 1
    });

    //クリックからフォーカス
    document.body.onclick = (e) => {
      const x = e.pageX;
      const y = e.pageY;

      const elementUnderMouse = document.elementFromPoint(x, y);
      if (elementUnderMouse.tagName == 'VIDEO') {
        this.focusThisVideo(elementUnderMouse.id);
      }
    };

    window.onpopstate = function () {
      console.log('webgazer is finish beacause browser back');
      webgazer.clearGazeListener().end();
      const gazeDotEl = document.getElementById('webgazerGazeDot');
      gazeDotEl.remove();
    };
  },
  watch: {
    isEnableGazeEstimating: function (isResumeButton) {
      console.log('isResumeButton', isResumeButton);
      if (isResumeButton) {
        if (this.isFirstGazeEstimating) {
          console.log('First begin gaze estimation');
          this.beginEstimateGaze();
          this.isFirstGazeEstimating = false;
          return;
        }

        this.resumeEstimateGaze();
      } else {
        this.pauseEstimateGaze();
      }
    }
  }
};
</script>

<style lang="scss">
body {
  -ms-overflow-style: none !important; /* IE, Edge 対応 */
  scrollbar-width: none !important; /* Firefox 対応 */
  &::-webkit-scrollbar {
    /* Chrome, Safari 対応 */
    display: none !important;
  }
}

.video {
  width: 100vw;
  padding: 10px 0;
  height: calc(100vh - 72px);
  display: flex;
  justify-content: space-around;
  align-items: center;
  flex-wrap: wrap;
  background-color: #d3be9a;
  overflow-y: scroll;
  -ms-overflow-style: none !important; /* IE, Edge 対応 */
  scrollbar-width: none !important; /* Firefox 対応 */
  &::-webkit-scrollbar {
    /* Chrome, Safari 対応 */
    display: none !important;
  }

  &-individual {
    width: 100%;
    height: 100%;
    border-radius: 80px;

    &-focus {
      border: solid 5px orange;
    }
  }
}

.status-bar {
  width: 100%;
  position: fixed;
  left: 0;
  bottom: 0;
}

.modal-open {
  & .modal-window {
    display: block;
  }
}

.modal-window {
  display: none;
  position: fixed;
  top: 0;
  left: 0;

  &-back {
    width: 100vw;
    height: 100vh;
    background-color: black;
    opacity: 0.7;
  }
  &-front {
    width: 50%;
    padding: 50px 0;
    position: fixed;
    text-align: center;
    max-width: 500px;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    background-color: #fff;
    border: solid 5px orange;
    border-radius: 8px;
    color: #000;
    font-size: 18px;
    font-weight: bold;

    &_btn-wrap {
      margin: 20px 0 0;
      text-align: center;
    }
  }
}
</style>
