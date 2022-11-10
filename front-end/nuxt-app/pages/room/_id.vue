<template>
  <section>
    <!-- TODO: 認識精度の確認用：エフェクトと連携したら削除 -->
    <!-- <div class=drinking>飲んだ秒数: {{ this.drinkingCount }}秒 / drinking: {{ this.accuracy.drinking }} / noDrinking: {{ this.accuracy.noDrinking }}</div> -->
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
        :drinkEstimatingFn="this.switchDrinkEstimating"
        :isEnableDrinkEstimating="this.isEnableDrinkEstimating"
        :restRoomStartFn="this.restRoomStart"
        :restRoomEndFn="this.restRoomEnd"
        :restRoomState="this.restRoomState"
      />
    </div>
    <!-- ビデオステータスバー -->

    <!-- スクリーンショットカウントダウンDialog-->
    <ScreenShotDialog
      :currentScreenShotCount="this.currentScreenShotCount"
      :isOpenScreenShotDialog="this.isOpenScreenShotDialog"
    />
    <!-- スクリーンショットカウントダウンDialog-->

    <!-- モーダルウィンドウ -->
    <section class="modal-window">
      <div class="modal-window-back"></div>
      <div class="modal-window-front">
        <h3>部屋に接続する</h3>
        <input type="text" placeholder="名前を入力" class="modal-window-front-input" id="inputName" />
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
import ScreenShotDialog from '~/components/presentational/organisms/screenShot/screenShotDialog';

import webgazer from 'webgazer';
import * as tf from '@tensorflow/tfjs';
import * as tmImage from '@teachablemachine/image';
import shutter from '~/assets/music/camera.mp3';

export default {
  components: {
    VideoState,
    Btn,
    AdjustWebgazerDialog,
    Video,
    Loader,
    ShareCard,
    ScreenShotDialog
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
      currentScreenShotCount: 3,
      isOpenScreenShotDialog: false,
      isEnableGazeEstimating: false,
      isEnableDrinkEstimating: false,
      isFirstGazeEstimating: true,
      elementUnderGazeCount: 0,
      myVideoStatus: true,
      myAudioStatus: true,
      focusThisVideoAllLiftCount: 0,
      roomMemberNumCheckIntervalFn: null,
      roomLeavingCheckTimeoutFn: null,
      isVisibleLoader: true,
      isConnectionRoom: false,
      // 飲み動作推定プロパティ
      model: null,
      webcam: null,
      labelContainer: null,
      maxPredictions: null,
      baseURL: 'https://teachablemachine.withgoogle.com/models/F2AWA0Eay/',
      drinkingCount: 0, // 飲んだ回数
      predictionCount: 0, // 推定結果の返却回数
      accuracy: { drinking: 0, noDrinking: 0 },
      isLoop: true,
      dialog: false,
      restRoomState: false,
      websocketNormalTermination: false
    };
  },

  methods: {
    //スクリーンショット起動
    captureImage() {
      const data = {
        type: 'SET_SCREEN_SHOT'
      };

      this.websocketConn.send(JSON.stringify(data));
    },

    setWebsocketEventListener: function (websocketConn) {
      websocketConn.onopen = function (e) {
        console.log('websocket connection');
      };
      websocketConn.onmessage = function (evt) {
        // フォーカスされた人以外の音量
        const unfocusedVolume = 0.15;

        // ======================================================================== //
        // add new member function
        // ======================================================================== //
        const addNewMemberRun = (data) => {
          // 新規メンバーを追加する(localStorage)
          for (let member in data.members) {
            localStorage.setItem(member, data.members[member].name);
          }
        };
        // ======================================================================== //

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
            // 誰もフォーカスしてない場合
            const videos = document.querySelectorAll('.video-individual');

            for (let video of videos) {
              //音量設定
              video.volume = 1;
            }
            return;
          }

          document.querySelector('#videomy-video').classList.add('video-individual-focus');

          for (let tgMemberData of focusMemberData) {
            let videoDom = document.getElementById(tgMemberData.peer_id);
            videoDom.classList.add('video-individual-focus');
            videoDom.volume = 1;
          }
        };
        const focusRun = (data, myVideoDomName) => {
          // フォーカス起動
          focusThisVideoAllLift();

          const memberDatas = data.focus_members;

          for (let memberData of memberDatas) {
            if (memberData.peer_id == myVideoDomName) {
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
          if (data.effect_member.peer_id == myVideoDomName) return;

          this.$refs.videoComponents.effectOthers(data.effect_member.type, data.effect_member.peer_id);
        };
        // ======================================================================== //

        // ======================================================================== //
        // screenshot function
        // ======================================================================== //
        const screenShotRun = () => {
          // エフェクト起動
          this.isOpenScreenShotDialog = true;

          const audio = new Audio(shutter);

          const countDown = () => {
            setTimeout(() => {
              if (this.currentScreenShotCount < 1) {
                this.isOpenScreenShotDialog = false;
                audio.play();
                html2canvas(document.querySelector('#capture')).then((canvas) => {
                  const link = document.createElement('a');
                  const number = Math.floor(Math.random() * 10000);
                  link.href = canvas.toDataURL();
                  link.download = `export_image_${number}.png`;
                  link.click();
                });
                return;
              }
              if (this.currentScreenShotCount > 0) {
                this.currentScreenShotCount = this.currentScreenShotCount - 1;
              }
              countDown();
            }, 1000);
          };

          countDown();
          this.currentScreenShotCount = 3;
        };
        // ======================================================================== //

        // ======================================================================== //
        // restRoom function
        // ======================================================================== //
        const restRoomRun = (data, myVideoDomName) => {
          // トイレ起動
          for (let member in data.members) {
            if (member == myVideoDomName) {
              this.$refs.videoComponents.restRoomOperationOnMySelf(data.members[member].isRestRoom);
              continue;
            }

            this.$refs.videoComponents.restRoomOperationOthers(member, data.members[member].isRestRoom);
          }
        };
        // ======================================================================== //

        const data = JSON.parse(evt.data);
        const myVideoDomName = document.querySelector('#videomy-video').getAttribute('name');

        console.log(data);

        // メンバー新規追加(自分含む)
        if (data.event_type == 'ADD_NEW_MEMBER') addNewMemberRun(data);

        // フォーカス
        if (data.event_type == 'SET_FOCUS' || data.event_type == 'DEL_ALL_FOCUS' || data.event_type == 'DEL_FOCUS')
          focusRun(data, myVideoDomName);

        // エフェクト
        if (data.event_type == 'SET_EFFECT') effectRun(data, myVideoDomName);

        // スクショ
        if (data.event_type == 'SET_SCREEN_SHOT') screenShotRun();

        // トイレ
        if (data.event_type == 'SET_REST_ROOM') restRoomRun(data, myVideoDomName);

        return;
      }.bind(this);
      websocketConn.onclose = function (evt) {
        console.log('websocket connection closed');

        if (!this.websocketNormalTermination) {
          console.log('websocketが異常終了したため再接続します');
          this.websocketConn = new WebSocket('wss://f-2205-server-chhumpv4gq-de.a.run.app/ws/' + this.$route.params.id);
          this.setWebsocketEventListener(this.websocketConn);
        }
      }.bind(this);
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
      if (document.querySelector('#inputName').value.length >= 10) {
        alert('最大文字数は10文字です');
        return;
      }

      this.isConnectionRoom = true;

      console.log('my peer id: ' + this.peer.id);
      const roomName = this.$route.params.id;
      const mediaConnection = this.peer.joinRoom(roomName, { mode: 'sfu', stream: this.localStream });
      this.setSkywayEventListener(mediaConnection);
      document.querySelector('body').classList.remove('modal-open');
      this.isVisibleSwitchButton = true;

      //クリックしたときのイベントを設定(フォーカス)
      this.setClickEvent();

      //人数制限チェック
      setTimeout(this.roomMemberNumCheck, 5000);
      this.roomMemberNumCheckIntervalFn = setInterval(this.roomMemberNumCheck, 60000);

      //nameを取得
      let name = document.querySelector('#inputName').value;
      if (!name) name = '未設定';

      const data = {
        type: 'ADD_NEW_MEMBER',
        info: {
          member: {
            name: name,
            peer_id: `video${this.peer.id}`
          }
        }
      };

      this.websocketConn.send(JSON.stringify(data));
    },

    roomLeaving: async function () {
      this.websocketNormalTermination = true;

      //ルーム退出
      console.log('room leaving start');

      //peerIDを破棄
      this.peer.destroy();

      //localStorageにあるmember情報を破棄
      localStorage.clear();

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
                from: `video${this.peer.id}`
              }
            }
          };

          console.log(data);

          this.websocketConn.send(data);
        } else {
          let response;
          response = await axios.delete('https://f-2205-server-chhumpv4gq-de.a.run.app/room/' + this.$route.params.id);
          console.log(response);
          response = await axios.delete('https://f-2205-server-chhumpv4gq-de.a.run.app/ws/' + this.$route.params.id);
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

    switchDrinkEstimating: function () {
      this.isEnableDrinkEstimating = !this.isEnableDrinkEstimating;
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
            peer_id: `video${this.peer.id}`,
            type: `${effectNumber}`
          }
        }
      };
      this.websocketConn.send(JSON.stringify(data));

      this.$refs.videoComponents.effectOnMySelf(`${effectNumber}`);
    },

    loaderOperation: function () {
      this.isVisibleLoader = !this.isVisibleLoader;
    },

    beginEstimateDrinking: function () {
      this.isLoop = true;
      this.initializeDrinkingModel();
    },

    endEstimateDrinking: function () {
      this.isLoop = false;
    },

    initializeDrinkingModel: async function () {
      const modelURL = this.baseURL + 'model.json';
      const metadataURL = this.baseURL + 'metadata.json';

      // modelとmetadataのロード
      this.model = await tmImage.load(modelURL, metadataURL);
      this.maxPredictions = this.model.getTotalClasses();
      window.requestAnimationFrame(this.loop);
    },

    loop: async function () {
      if (this.isLoop) {
        await this.predictDrinking();
        window.requestAnimationFrame(this.loop);
      }
    },

    predictDrinking: async function () {
      // 推定結果
      const prediction = await this.model.predict(document.getElementById('videomy-video'));
      this.accuracy.drinking = prediction[0].probability.toFixed(2);
      this.accuracy.noDrinking = prediction[1].probability.toFixed(2);

      // Drinking class の精度が8割以上の時、カウントを行う
      if (this.accuracy.drinking >= 0.8) {
        this.predictionCount += 1;

        // 数ミリ秒単位でカウントしているため，数回カウントで制御
        if (this.predictionCount > 100) {
          // 4秒程度
          this.effectFn('3'); // 後ほど4に変更
          this.drinkingCount += 1; // TODO: 廃止予定
          this.predictionCount = 0;
        }
      }
    },

    setClickEvent: function () {
      document.body.onclick = (e) => {
        const x = e.pageX;
        const y = e.pageY;

        const elementUnderMouse = document.elementFromPoint(x, y);

        if (elementUnderMouse.tagName == 'VIDEO' && elementUnderMouse.id != 'videomy-video')
          this.focusThisVideo(elementUnderMouse.id);

        if (elementUnderMouse.id == 'video-wrap' || elementUnderMouse.id == 'video-box') this.focusThisVideoAllLift();
      };
    },

    restRoomStart: function () {
      // トイレ機能開始
      const data = {
        type: 'SET_REST_ROOM',
        info: {
          rest: {
            peer_id: `video${this.peer.id}`,
            is_rest_room: true
          }
        }
      };
      this.websocketConn.send(JSON.stringify(data));

      this.videoMute();
      this.audioMute();

      this.restRoomState = true;
    },

    restRoomEnd: function () {
      // トイレ機能終了
      const data = {
        type: 'SET_REST_ROOM',
        info: {
          rest: {
            peer_id: `video${this.peer.id}`,
            is_rest_room: false
          }
        }
      };
      this.websocketConn.send(JSON.stringify(data));

      this.videoMute();
      this.audioMute();

      this.restRoomState = false;
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

    //ルーム接続時間制限(50分)
    this.roomLeavingCheckTimeoutFn = setTimeout(this.roomLeaving, 3000000);

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

      if (elementUnderMouse.tagName == 'VIDEO' && elementUnderMouse.id != 'videomy-video')
        this.focusThisVideo(elementUnderMouse.id);

      if (elementUnderMouse.id == 'videomy-video') this.focusThisVideoAllLift();

      if (elementUnderMouse.id == 'video-wrap') this.focusThisVideoAllLift();
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
    },

    isEnableDrinkEstimating: function (isResumeButton) {
      console.log('isResumeDrinkingButton', isResumeButton);
      if (isResumeButton) {
        this.beginEstimateDrinking();
        console.log('beginEstimateDrinking');
      } else {
        this.endEstimateDrinking();
        console.log('endEstimateDrinking');
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

    &-input {
      margin: 40px 0 10px;
      border-bottom: solid 2px #ccc;
      background-color: #fff;
      outline: none !important;
      transition: all 0.4s;
      &:hover {
        border-bottom: solid 2px orange;
      }
      &:focus {
        border-bottom: solid 2px orange;
      }
    }

    &_btn-wrap {
      margin: 20px 0 0;
      text-align: center;
    }
  }
}
</style>
