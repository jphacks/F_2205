<template>
  <section id="video-wrap" class="video">
    <div class="video-line">
      <video id="my-video" class="video-individual" autoplay muted playsinline></video>
    </div>
    <VideoEffect ref="effectComponentsMyVideo" videoId="my-video" />
  </section>
</template>

<script>
import Vue from 'vue/dist/vue.esm.js';
import VideoEffect from '~/components/presentational/organisms/videoEffect';

export default {
  props: ['roomMemberNum'],
  components: {
    VideoEffect
  },

  methods: {
    addVideo: function (stream, roomMemberNum) {
      const videoLineDoms = document.querySelectorAll('.video-line');

      const videoDom = document.createElement('video');
      videoDom.setAttribute('id', stream.peerId);
      videoDom.classList.add('video-individual');
      videoDom.srcObject = stream;
      videoDom.play();

      //append
      let divDom;
      if (roomMemberNum <= 3) {
        videoLineDoms[0].append(videoDom);

        this.addEffectComponents(stream.peerId);
        return;
      }

      if (roomMemberNum > 3 && roomMemberNum <= 6) {
        if (roomMemberNum == 4) {
          divDom = document.createElement('div');
          divDom.classList.add('video-line');
          divDom.append(videoDom);
          document.getElementById('video-wrap').append(divDom);

          this.addEffectComponents(stream.peerId);
          return;
        }

        videoLineDoms[1].append(videoDom);
        this.addEffectComponents(stream.peerId);
        return;
      }

      if (roomMemberNum > 6 && roomMemberNum <= 9) {
        if (roomMemberNum == 7) {
          divDom = document.createElement('div');
          divDom.classList.add('video-line');
          divDom.append(videoDom);
          document.getElementById('video-wrap').append(divDom);

          this.addEffectComponents(stream.peerId);
          return;
        }

        videoLineDoms[2].append(videoDom);
        this.addEffectComponents(stream.peerId);
        return;
      }

      if (roomMemberNum > 9 && roomMemberNum <= 12) {
        if (roomMemberNum == 10) {
          divDom = document.createElement('div');
          divDom.classList.add('video-line');
          divDom.append(videoDom);
          document.getElementById('video-wrap').append(divDom);

          this.addEffectComponents(stream.peerId);
          return;
        }

        videoLineDoms[3].append(videoDom);
        this.addEffectComponents(stream.peerId);
        return;
      }

      if (roomMemberNum > 12 && roomMemberNum <= 15) {
        if (roomMemberNum == 13) {
          divDom = document.createElement('div');
          divDom.classList.add('video-line');
          divDom.append(videoDom);
          document.getElementById('video-wrap').append(divDom);

          this.addEffectComponents(stream.peerId);
          return;
        }

        videoLineDoms[4].append(videoDom);
        this.addEffectComponents(stream.peerId);
        return;
      }

      if (roomMemberNum > 15 && roomMemberNum <= 18) {
        if (roomMemberNum == 16) {
          divDom = document.createElement('div');
          divDom.classList.add('video-line');
          divDom.append(videoDom);
          document.getElementById('video-wrap').append(divDom);

          this.addEffectComponents(stream.peerId);
          return;
        }

        videoLineDoms[5].append(videoDom);
        this.addEffectComponents(stream.peerId);
        return;
      }
    },

    removeVideo: function (peerId, roomMemberNum) {
      console.log(roomMemberNum);

      const videoDom = document.getElementById(peerId);
      videoDom.remove();

      if (roomMemberNum % 3 == 0) {
        const videoLineDom = document.querySelectorAll('.video-line');
        console.log(videoLineDom);
        videoLineDom[videoLineDom.length - 1].remove();
      }
    },

    addEffectComponents: function (videoId) {
      const ComponentClass = Vue.extend(VideoEffect);
      const instance = new ComponentClass({
        propsData: {
          ref: `effectComponents${videoId}`,
          videoId: videoId
        }
      });
      instance.$mount();

      document.querySelector('#video-wrap').append(instance.$el);
    },

    effectOthers: function (effectNumber, videoId) {
      //自分以外のビデオにエフェクトを追加する(websocketConn.onmessageから呼ばれる)
      const tgRef = 'effectComponents' + videoId;
      this.$refs.eval(tgRef).start(effectNumber);
    },

    effectOnMySelf: function (effectNumber) {
      //自分の画像にエフェクトを追加する
      this.$refs.effectComponentsMyVideo.start(effectNumber);
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
  height: calc(100vh - 72px);
  background-color: #d3be9a;
  overflow-y: scroll;
  -ms-overflow-style: none !important; /* IE, Edge 対応 */
  scrollbar-width: none !important; /* Firefox 対応 */
  &::-webkit-scrollbar {
    /* Chrome, Safari 対応 */
    display: none !important;
  }

  &-line {
    width: 100vw;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  &-individual {
    flex: 1;
    margin: 10px;
    border-radius: 20px;
    max-width: 480px;

    &-focus {
      border: solid 5px orange;
    }
  }
}
</style>
