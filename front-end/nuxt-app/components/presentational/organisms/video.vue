<template>
  <section id="video-wrap" class="video">
    <!-- ビデオレイヤー -->
    <div class="video-box">
      <video id="videomy-video" class="video-individual" autoplay muted playsinline></video>
    </div>
    <!-- ビデオレイヤー -->

    <!-- 名前レイヤー -->
    <div v-for="(peerId, index) of peerIds" :key="`${peerId}Name-${index}`">
      <VideoName v-if="peerId != undefined" :ref="`nameComponentsvideo${peerId}`" :videoId="peerId" />
    </div>
    <!-- 名前レイヤー -->

    <!-- エフェクトレイヤー -->
    <div v-for="(peerId, index) of peerIds" :key="`${peerId}Effect-${index}`">
      <VideoEffect v-if="peerId != undefined" :ref="`effectComponentsvideo${peerId}`" :videoId="peerId" />
    </div>
    <!-- エフェクトレイヤー -->
  </section>
</template>

<script>
import VideoEffect from '~/components/presentational/organisms/videoEffect';
import VideoName from '~/components/presentational/organisms/videoName';

export default {
  props: ['roomMemberNum'],
  components: {
    VideoEffect,
    VideoName
  },

  data() {
    return {
      peerIds: ['my-video']
    };
  },

  methods: {
    addVideo: function (stream, roomMemberNum) {
      this.videoNameAllRePositionWaitDisplay();

      const videoBoxDom = document.querySelector('.video-box');

      const videoDom = document.createElement('video');

      videoDom.setAttribute('id', `video${stream.peerId}`);
      videoDom.classList.add('video-individual');
      videoDom.srcObject = stream;
      videoDom.play();

      // append
      videoBoxDom.append(videoDom);
      this.addComponents(stream.peerId);

      // resize
      this.videoResize(roomMemberNum);
      // reposition
      setTimeout(this.videoNameAllRePosition, 5000);

      return;
    },

    removeVideo: function (peerId, roomMemberNum) {
      console.log(roomMemberNum);

      // ビデオを削除
      const videoDom = document.getElementById(`video${peerId}`);
      videoDom.remove();

      // ビデオネームを削除
      const tgRef = 'nameComponents' + 'video' + peerId;
      this.$refs[tgRef][0].remove();

      // peerIds配列から削除
      const deleteTgIndex = this.peerIds.indexOf(peerId);
      delete this.peerIds[deleteTgIndex];

      // resize
      this.videoResize(roomMemberNum);
      // reposition
      this.videoNameAllRePosition();
    },

    videoResize: function (roomMemberNum) {
      console.log('resize video size');

      let windowSizeW;
      const windowSizeH = document.body.clientHeight;
      document.body.clientWidth > 1400 ? (windowSizeW = 1400) : (windowSizeW = document.body.clientWidth);
      const DisplayVideoH = windowSizeH - 79;

      const videoDoms = document.querySelectorAll('.video-individual');

      // 初期化
      document.querySelector('#video-wrap').style.padding = '10px 0';
      for (let videoDom of videoDoms) {
        videoDom.style.marginTop = '10px';
        videoDom.style.marginBottom = '10px';
      }

      let line;
      let column;
      let margenPx;
      let margenPxW;
      let margenPxH;
      let videoSizeH;
      let videoSizeW;
      let videoSizeWFromH;
      let videoSizeWFromW;
      let videoSizeWAverage;
      let decreaseConstant;

      if (roomMemberNum >= 16) {
        column = 5;
        margenPx = 10 * (column * 2);

        //ビデオの幅
        videoSizeW = (windowSizeW - margenPx) / column;

        document.querySelector('#video-wrap').style.padding = '1px 0';

        for (let videoDom of videoDoms) {
          videoDom.style.width = videoSizeW + 'px';
          videoDom.style.marginTop = '1px';
          videoDom.style.marginBottom = '1px';
        }
        return;
      }

      if (roomMemberNum >= 13) {
        column = 5;
        margenPx = 10 * (column * 2);

        //ビデオの幅
        videoSizeW = (windowSizeW - margenPx) / column;

        for (let videoDom of videoDoms) {
          videoDom.style.width = videoSizeW + 'px';
          videoDom.style.marginTop = '4px';
          videoDom.style.marginBottom = '4px';
        }
        return;
      }

      if (roomMemberNum >= 10) {
        column = 4;
        margenPx = 10 * (column * 2);

        //ビデオの幅
        videoSizeW = (windowSizeW - margenPx) / column;

        for (let videoDom of videoDoms) {
          videoDom.style.width = videoSizeW + 'px';
          videoDom.style.marginTop = '4px';
          videoDom.style.marginBottom = '4px';
        }
        return;
      }

      if (roomMemberNum >= 7) {
        line = 3;
        column = 3;
        decreaseConstant = 20;
        margenPxH = 10 * (line * 2);
        margenPxW = 10 * (column * 2);

        // ビデオの欲しい高さ
        videoSizeH = DisplayVideoH / line - margenPxH;

        // ビデオの幅(高さ視点)
        videoSizeWFromH = Math.round(videoSizeH * 1.3);

        // ビデオの幅(幅視点)
        videoSizeWFromW = (windowSizeW - margenPxW) / 3;

        // 平均ビデオ幅(少しマイナス)
        videoSizeWAverage = (videoSizeWFromH + videoSizeWFromW) / 2 - decreaseConstant;

        for (let videoDom of videoDoms) {
          videoDom.style.width = videoSizeWAverage + 'px';
          videoDom.style.marginTop = '4px';
          videoDom.style.marginBottom = '4px';
        }

        return;
      }

      if (roomMemberNum >= 5) {
        column = 3;
        margenPx = 10 * (column * 2);

        //ビデオの幅
        videoSizeW = (windowSizeW - margenPx) / column;

        for (let videoDom of videoDoms) {
          videoDom.style.width = videoSizeW + 'px';
        }
        return;
      }

      if (roomMemberNum >= 3) {
        line = 2;
        margenPx = 10 * (line * 2);

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
        column = 2;
        margenPx = 10 * (column * 2);

        //ビデオの幅
        videoSizeW = (windowSizeW - margenPx) / column;

        for (let videoDom of videoDoms) {
          videoDom.style.width = videoSizeW + 'px';
          videoDom.style.marginTop = '4px';
          videoDom.style.marginBottom = '4px';
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

    addComponents: function (videoId) {
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
    },

    videoNameAllRePositionWaitDisplay: function () {
      let tgRef;
      for (let peerId of this.peerIds) {
        if (peerId == undefined) continue;
        tgRef = 'nameComponentsvideo' + peerId;
        this.$refs[tgRef][0].setPositionWaitDisplay();
      }
    },
    videoNameAllRePosition: function () {
      let tgRef;
      for (let peerId of this.peerIds) {
        if (peerId == undefined) continue;
        tgRef = 'nameComponentsvideo' + peerId;
        this.$refs[tgRef][0].setPosition();
      }
    }
  },

  mounted: function () {
    // resizeイベント登録
    window.addEventListener('resize', () => {
      console.log('resize');

      // videoNameのリサイズ
      this.videoNameAllRePosition();

      // videoのリサイズ
      this.videoResize(this.roomMemberNum);
    });
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
    max-width: 1400px;
    margin: 0 auto;

    * {
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
