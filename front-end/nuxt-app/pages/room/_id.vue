<template>
  <section>
    <div id="capture">
      <Video ref="videoComponents" :roomMemberNum="this.roomMemberNum" />
    </div>
    <!-- シェアカード -->
    <ShareCard />
    <!-- シェアカード -->

    <!-- ビデオステータスバー -->
    <div class="status-bar">
      <VideoState
        :leavingFn="this.roomLeaving"
        :gazeEstimatingFn="this.swtichEstimateGaze"
        :isEnableGazeEstimating="this.isEnableGazeEstimating"
        :focusThisVideoAllLiftFn="this.focusThisVideoAllLift"
        :isOpenAdjustWebGazerDialog="this.isOpenAdjustWebGazerDialog"
        :handleAdjustWebGazer="this.handleAdjustWebGazer"
        :videoMuteFn="this.videoMute"
        :audioMuteFn="this.audioMute"
        :myVideoStatus="this.myVideoStatus"
        :myAudioStatus="this.myAudioStatus"
        :captureImage="this.captureImage"
        :effectFn="this.effectFn"
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

    <!-- ローダー -->
    <Loader v-if="isVisibleLoader" @loaderOperationFn="loaderOperation" :autoDelete="true" />
    <!-- ローダー -->
  </section>
</template>

<script>
import Peer from 'skyway-js';
import axios from 'axios';
import html2canvas from 'html2canvas';

import VideoState from '~/components/presentational/organisms/videoState';
import Btn from '~/components/presentational/atoms/btn';
import AdjustWebgazerDialog from '~/components/presentational/organisms/adjustWebgazerDialog';
import Video from '~/components/presentational/organisms/video';
import Loader from '~/components/presentational/organisms/loader';
import ShareCard from '~/components/presentational/organisms/shareCard';

import webgazer from 'webgazer';

export default {
  components: {
    VideoState,
    Btn,
    AdjustWebgazerDialog,
    Video,
    Loader,
    ShareCard
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
      isOpenAdjustWebGazerDialog: false,
      isEnableGazeEstimating: false,
      isFirstGazeEstimating: true,
      elementUnderGazeCount: 0,
      myVideoStatus: true,
      myAudioStatus: true,
      focusThisVideoAllLiftCount: 0,
      roomMemberNumCheckIntervalFn: null,
      roomLeavingCheckTimeoutFn: null,
      isVisibleLoader: true,
      isConnectionRoom: false
    };
  },

  methods: {
    //スクリーンショット起動
    captureImage() {
      html2canvas(document.querySelector('#capture')).then((canvas) => {
        const link = document.createElement('a');
        link.href = canvas.toDataURL();
        link.download = `export_image.png`;
        link.click();
      });
    },
    setWebsocketEventListener: function (websocketConn) {
      websocketConn.onopen = function (e) {
        console.log('websocket connection');
      };
      websocketConn.onmessage = function (evt) {
        // フォーカスされた人以外の音量
        const unfocusedVolume = 0.15;

        // ======================================================================== //
        // focus function
        // ======================================================================== //
        const focusThisVideoAllLift = () => {
          // フォーカス全解除&音量下げ

          const videos = document.querySelectorAll('.video-individual');

          for (let video of videos) {
            //音量設定
            video.volume = unfocusedVolume;
            video.classList.remove('video-individual-focus');
          }
        };
        const doFocus = (focusMemberData) => {
          // 引数のメンバーをフォーカス

          if (focusMemberData.length == 0) {
            const videos = document.querySelectorAll('.video-individual');

            for (let video of videos) {
              //音量設定
              video.volume = 1;
            }
            return;
          }

          document.querySelector('#videomy-video').classList.add('video-individual-focus');

          for (let tgMemberData of focusMemberData) {
            let videoDom = document.getElementById(tgMemberData.name);
            videoDom.classList.add('video-individual-focus');
            videoDom.volume = 1;
          }
        };

        const focusRun = (data, myVideoDomName) => {
          // フォーカス起動
          focusThisVideoAllLift();

          const memberDatas = data.focus_members;

          for (let memberData of memberDatas) {
            if (memberData.name == myVideoDomName) {
              //引数のメンバーをフォーカスしてそれ以外を除外
              doFocus(memberData.connects);
            }
          }
        };
        // ======================================================================== //

        // ======================================================================== //
        // effect function
        // ======================================================================== //
        const effectRun = (data, myVideoDomName) => {
          // エフェクト起動
          if (data.effect_member.name == myVideoDomName) return;

          this.$refs.videoComponents.effectOthers(data.effect_member.type, data.effect_member.name);
        };
        // ======================================================================== //

        const data = JSON.parse(evt.data);
        const myVideoDomName = document.querySelector('#videomy-video').getAttribute('name');

        // フォーカス
        if (data.event_type == 'SET_FOCUS' || data.event_type == 'DEL_ALL_FOCUS' || data.event_type == 'DEL_FOCUS')
          focusRun(data, myVideoDomName);

        // エフェクト
        if (data.event_type == 'SET_EFFECT') effectRun(data, myVideoDomName);

        return;
      }.bind(this);
      websocketConn.onclose = function (evt) {
        console.log('websocket connection closed');
      };
      websocketConn.onerror = function (evt) {
        console.log('websocket error: ' + evt);
      };
    },
    setSkywayEventListener: function (mediaConnection) {
      mediaConnection.on('stream', (stream) => {
        //ルームメンバー人数追加
        this.roomMemberNum++;

        //video要素にカメラ映像をセットして再生
        this.$refs.videoComponents.addVideo(stream, this.roomMemberNum);
      });

      mediaConnection.on('open', () => {
        document.querySelector('#videomy-video').setAttribute('name', `video${this.peer.id}`);
      });

      //ルームメンバーが退出したときに発火
      mediaConnection.on('peerLeave', (peerId) => {
        //ルームメンバー人数減少
        this.roomMemberNum--;

        //ビデオの削除
        this.$refs.videoComponents.removeVideo(peerId, this.roomMemberNum);
      });

      //何らかのエラーが発生した場合に発火
      this.peer.on('error', (err) => {
        alert(err.message);
      });
    },

    roomConnection: function () {
      //ルーム作成 or ルーム参加
      this.isConnectionRoom = true;

      console.log('my peer id: ' + this.peer.id);
      const roomName = this.$route.params.id;
      const mediaConnection = this.peer.joinRoom(roomName, { mode: 'sfu', stream: this.localStream });
      this.setSkywayEventListener(mediaConnection);
      document.querySelector('body').classList.remove('modal-open');
      this.isVisibleSwitchButton = true;

      //人数制限チェック
      setTimeout(this.roomMemberNumCheck, 5000);
      this.roomMemberNumCheckIntervalFn = setInterval(this.roomMemberNumCheck, 60000);

      const data = {
        type: 'NEW_MEMBER',
        info: {
          focus: {
            from: `video${this.peer.id}`
          }
        }
      };

      this.websocketConn.send(JSON.stringify(data));
    },

    roomLeaving: async function () {
      //ルーム退出
      console.log('room leaving start');

      //peerIDを破棄
      this.peer.destroy();

      //人数チェック処理を解除(Interval)
      clearInterval(this.roomMemberNumCheckIntervalFn);
      //強制退出処理を解除(Timeout)
      clearTimeout(this.roomLeavingCheckTimeoutFn);

      //websocket そのユーザーの持っている接続状態を解除する
      if (this.isConnectionRoom) {
        if (this.roomMemberNum != 1) {
          const data = {
            type: 'DEL_ALL_FOCUS',
            info: {
              focus: {
                from: `video${this.peer}`
              }
            }
          };

          this.websocketConn.send(data);
        } else {
          const response = await axios.delete(
            'https://f-2205-server-chhumpv4gq-de.a.run.app/ws/' + this.$route.params.id
          );
          console.log(response);
        }
      }

      this.isConnectionRoom = false;

      //websocketの接続を切断(正常終了)
      this.websocketConn.close(1000, 'normal amputation websocket');

      //end EstimateGaze
      if (this.isEnableGazeEstimating) {
        this.endEstimateGaze();
        this.isVisibleSwitchButton = false;
      }

      alert('退出しました');

      //リダイレクト
      this.$router.push('/');
    },

    beginEstimateGaze: function () {
      //視線からフォーカス
      webgazer
        .showVideo(false)
        .showPredictionPoints(true)
        .setRegression('ridge')
        .setTracker('clmtrackr')
        .applyKalmanFilter(true)
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
            this.focusThisVideoAllLiftCount = 0;
            // TODO: 試験的にカウントを10以上に設定, 後ほど適切な値・実装方法に変える
            if (this.elementUnderGazeCount > 10) {
              console.log('elementUnderGazeCount is 10 count');
              this.focusThisVideoLineOfSight(elementUnderGaze.id);
            }
          } else {
            this.focusThisVideoAllLiftCount++;
            this.elementUnderGazeCount = 0;
            //フォーカス全外し(この関数を呼ぶことでサーバー側にリクエスト飛ぶ)
            if (this.focusThisVideoAllLiftCount > 10) {
              console.log('focusThisVideoAllLift is 10 count');
              this.focusThisVideoAllLift();
            }
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
      if (id == 'videomy-video') return;

      const className = document.getElementById(id).className;
      if (className == 'video-individual') {
        //websocket ユーザー同士を接続状態にする
        const data = {
          type: 'SET_FOCUS',
          info: {
            focus: {
              from: `video${this.peer.id}`,
              to: `${id}`
            }
          }
        };
        this.websocketConn.send(JSON.stringify(data));
      } else {
        //websocket ユーザー同士を接続状態にする
        const data = {
          type: 'DEL_FOCUS',
          info: {
            focus: {
              from: `video${this.peer.id}`,
              to: `${id}`
            }
          }
        };
        this.websocketConn.send(JSON.stringify(data));
      }
    },
    focusThisVideoLineOfSight: function (id) {
      //ビデオをフォーカスする(自分のビデオ以外)(視線で)
      if (id == 'videomy-video') return;

      //websocket ユーザー同士を接続状態にする
      const data = {
        type: 'SET_FOCUS',
        info: {
          focus: {
            from: `video${this.peer.id}`,
            to: `${id}`
          }
        }
      };
      this.websocketConn.send(JSON.stringify(data));
    },
    focusThisVideoAllLift: function () {
      //フォーカスを全解除
      const data = {
        type: 'DEL_ALL_FOCUS',
        info: {
          focus: {
            from: `video${this.peer.id}`
          }
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
      console.log('There are currently ' + this.roomMemberNum + ' people in the room');
    },
    handleAdjustWebGazer: function () {
      //視線調整用Dialog開閉
      this.isOpenAdjustWebGazerDialog = !this.isOpenAdjustWebGazerDialog;
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
    },

    effectFn: function (effectNumber) {
      //自分の画像にエフェクトを追加する(エフェクト作動)
      const data = {
        type: 'SET_EFFECT',
        info: {
          effect: {
            name: `video${this.peer.id}`,
            type: `${effectNumber}`
          }
        }
      };
      this.websocketConn.send(JSON.stringify(data));

      this.$refs.videoComponents.effectOnMySelf(`${effectNumber}`);
    },

    loaderOperation: function () {
      this.isVisibleLoader = !this.isVisibleLoader;
    }
  },

  mounted: async function () {
    if (!this.$route.params.id) {
      alert('部屋番号が入力されていません');
      this.$router.push('/room/prepare');
    }

    //WebSocketで接続
    this.websocketConn = new WebSocket('wss://f-2205-server-chhumpv4gq-de.a.run.app/ws/' + this.$route.params.id);
    this.setWebsocketEventListener(this.websocketConn);

    //ルーム接続時間制限(5分)
    this.roomLeavingCheckTimeoutFn = setTimeout(this.roomLeaving, 300000);

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
      // カメラ映像取得(自)
      const stream = await navigator.mediaDevices.getUserMedia(constraints);
      document.getElementById('videomy-video').srcObject = stream;
      this.localStream = stream;
    } catch (error) {
      console.log(error);
    }

    //モーダルウィンドウを表示
    document.querySelector('body').classList.add('modal-open');

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
