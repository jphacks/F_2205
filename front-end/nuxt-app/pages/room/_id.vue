<template>
  <section>
    <section id="video-wrap" class="video">
      <video id="my-video" class="video-individual" autoplay muted playsinline></video>
    </section>

    <Btn text="接続" color="orange" :clickedfn="this.roomConnection" />

    <!-- ビデオステータスバー -->
    <VideoState :leavingFn="this.roomLeaving" />
    <!-- ビデオステータスバー -->
  </section>
</template>

<script>
import Peer from 'skyway-js';

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
      localStream: null,
      peer: null
    };
  },

  methods: {
    setSkywayEventListener: function (mediaConnection) {
      mediaConnection.on('stream', (stream) => {
        // video要素にカメラ映像をセットして再生
        this.addVideo(stream);
        const videoElm = document.getElementById(stream.id);
        videoElm.srcObject = stream;
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
      console.log(this.peer);
      const roomName = this.$route.params.id;
      const mediaConnection = this.peer.joinRoom(roomName, { mode: 'sfu', stream: this.localStream });
      this.setSkywayEventListener(mediaConnection);
    },

    roomLeaving: function () {
      //ルーム退出
      this.peer.destroy();
      alert('退出しました');
      this.$router.push('/room/prepare');
    },

    addVideo: function (stream) {
      const videoDom = document.createElement('video');
      videoDom.setAttribute('id', stream.peerId);
      videoDom.classList.add('video-individual');
      videoDom.srcObject = stream;
      videoDom.play();
      document.getElementById('video-wrap').append(videoDom);
    },
    removeVideo: function (peerId) {
      console.log(peerId);
      const videoDom = document.getElementById(peerId);
      videoDom.remove();
    },

    focusThisVideo: function (id) {
      //クリックされたビデオをフォーカスする(自分のビデオ以外)
      const videoId = id;
      if (videoId == 'my-video') {
        console.log('my-video');
        return;
      }

      const videoDom = document.getElementById(videoId);
      const myVideoDom = document.getElementById('my-video');
      videoDom.classList.add('video-individual-focus');
      myVideoDom.classList.add('video-individual-focus');
      videoDom.volume = 0.1;
    }
  },

  mounted: async function () {
    if (!this.$route.params.id) {
      alert('部屋番号が入力されていません');
      this.$router.push('/room/prepare');
    }

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
      debug: 3
    });

    //マウスのクリックした場所のエレメントを取得
    document.body.onclick = (e) => {
      const x = e.pageX;
      const y = e.pageY;

      const elementUnderMouse = document.elementFromPoint(x, y);

      if (elementUnderMouse.tagName == 'VIDEO') {
        this.focusThisVideo(elementUnderMouse.id);
      }
    };

    //視線移動を行った場所のエレメントを取得
    webgazer
      .showVideo(false)
      .showPredictionPoints(true)
      .setGazeListener((gaze, clock) => {
        const x = gaze.x;
        const y = gaze.y;

        const elementUnderMouse = document.elementFromPoint(x, y);

        if (elementUnderMouse.tagName == 'VIDEO') {
          this.focusThisVideo(elementUnderMouse.id);
        }
      })
      .begin();
  }
};
</script>

<style lang="scss">
.video {
  display: flex;
  justify-content: space-around;
  align-items: center;
  flex-wrap: wrap;

  &-individual {
    width: 45%;

    &-focus {
      border: solid 3px red;
    }
  }
}
</style>