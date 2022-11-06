<template>
  <section id="video-wrap" class="video">
    <div class="video-box">
      <video id="videomy-video" class="video-individual" autoplay muted playsinline></video>
    </div>
    <div v-for="peerId of peerIds" :key="peerId">
      <VideoEffect :ref="`effectComponentsvideo${peerId}`" :videoId="peerId" />
    </div>
  </section>
</template>

<script>
import VideoEffect from '~/components/presentational/organisms/videoEffect';

export default {
  props: ['roomMemberNum'],
  components: {
    VideoEffect
  },

  data() {
    return {
      peerIds: ['my-video']
    };
  },

  methods: {
    addVideo: function (stream, roomMemberNum) {
      const videoBoxDom = document.querySelector('.video-box');

      const videoDom = document.createElement('video');

      videoDom.setAttribute('id', `video${stream.peerId}`);
      videoDom.classList.add('video-individual');
      videoDom.srcObject = stream;
      videoDom.play();

      // append
      videoBoxDom.append(videoDom);

      // resize
      this.videoResize(roomMemberNum);

      return;
    },

    removeVideo: function (peerId, roomMemberNum) {
      console.log(roomMemberNum);

      const videoDom = document.getElementById(`video${peerId}`);
      videoDom.remove();

      if (roomMemberNum % 3 == 0) {
        const videoLineDom = document.querySelectorAll('.video-line');
        console.log(videoLineDom);
        videoLineDom[videoLineDom.length - 1].remove();
      }

      // resize
      this.videoResize(roomMemberNum);
    },

    videoResize: function (roomMemberNum) {
      console.log('resize video size');

      const windowSizeH = document.body.clientHeight;
      const windowSizeW = document.body.clientWidth;
      const DisplayVideoH = windowSizeH - 79;

      const videoDoms = document.querySelectorAll('.video-individual');

      let line;
      let margenPx;
      let videoSizeH;
      let videoSizeW;

      if (roomMemberNum >= 16) {
        line = 6;
        margenPx = 10 * (line + 1);

        //ビデオの欲しい高さ
        videoSizeH = DisplayVideoH / line - margenPx;

        //ビデオの幅
        videoSizeW = Math.round(videoSizeH * 1.3);

        for (let videoDom of videoDoms) {
          videoDom.style.maxWidth = videoSizeW + 'px';
        }

        return;
      }

      if (roomMemberNum >= 13) {
        line = 5;
        margenPx = 10 * (line + 1);

        //ビデオの欲しい高さ
        videoSizeH = DisplayVideoH / line - margenPx;

        //ビデオの幅
        videoSizeW = Math.round(videoSizeH * 1.3);

        for (let videoDom of videoDoms) {
          videoDom.style.maxWidth = videoSizeW + 'px';
        }

        return;
      }

      if (roomMemberNum >= 10) {
        line = 4;
        margenPx = 10 * (line + 1);

        //ビデオの欲しい高さ
        videoSizeH = DisplayVideoH / line - margenPx;

        //ビデオの幅
        videoSizeW = Math.round(videoSizeH * 1.3);

        for (let videoDom of videoDoms) {
          videoDom.style.maxWidth = videoSizeW + 'px';
        }

        return;
      }

      if (roomMemberNum >= 7) {
        line = 3;
        margenPx = 10 * (line + 1);

        //ビデオの欲しい高さ
        videoSizeH = DisplayVideoH / line - margenPx;

        //ビデオの幅
        videoSizeW = Math.round(videoSizeH * 1.3);

        for (let videoDom of videoDoms) {
          videoDom.style.maxWidth = videoSizeW + 'px';
        }

        return;
      }

      if (roomMemberNum >= 3) {
        // 4人の場合
        line = 2;
        margenPx = 10 * (line + 1);

        //ビデオの欲しい高さ
        videoSizeH = DisplayVideoH / line - margenPx;

        //ビデオの幅
        videoSizeW = Math.round(videoSizeH * 1.3);

        for (let videoDom of videoDoms) {
          videoDom.style.width = videoSizeW + 'px';
        }
        return;
      }

      if (roomMemberNum == 2) {
        // 2人の場合

        //ビデオの欲しい高さ
        videoSizeH = DisplayVideoH - margenPx;

        //ビデオの幅
        videoSizeW = Math.round(videoSizeH * 1.3);

        for (let videoDom of videoDoms) {
          videoDom.style.width = videoSizeW + 'px';
        }
        return;
      }

      // デフォルト
      let defaultVideoSize = 480;
      for (let videoDom of videoDoms) {
        videoDom.style.width = defaultVideoSize + 'px';
      }

      return;
    },

    addEffectComponents: function (videoId) {
      this.peerIds.push(videoId);
    },

    effectOthers: function (effectNumber, videoDomId) {
      //自分以外のビデオにエフェクトを追加する(websocketConn.onmessageから呼ばれる)
      const tgRef = 'effectComponents' + videoDomId;
      this.$refs[tgRef][0].start(effectNumber);
    },

    effectOnMySelf: function (effectNumber) {
      //自分の画像にエフェクトを追加する
      this.$refs['effectComponentsvideomy-video'][0].start(effectNumber);
    }
  }
};
</script>

<style lang="scss">
.video {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  width: 100vw;
  padding: 10px 0;
  height: calc(100vh - 79px);
  background-color: #d3be9a;
  overflow-y: scroll;
  -ms-overflow-style: none !important; /* IE, Edge 対応 */
  scrollbar-width: none !important; /* Firefox 対応 */
  &::-webkit-scrollbar {
    /* Chrome, Safari 対応 */
    display: none !important;
  }

  &-box {
    width: 100vw;
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    justify-content: center;

    * {
      // flex: 1;
      margin: 10px;
      border-radius: 20px;
      max-width: 620px;
      width: 100%;
    }

    .video-individual-focus {
      border: solid 5px orange;
    }
  }

  &-individual {
  }
}
</style>
