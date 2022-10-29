<template>
  <section id="video-wrap" class="video">
    <div class="video-line">
      <video id="my-video" class="video-individual" autoplay muted playsinline></video>
    </div>
  </section>
</template>

<script>
export default {
  props: ['roomMemberNum'],

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
        return;
      }

      if (roomMemberNum > 3 && roomMemberNum >= 6) {
        if (roomMemberNum == 4) {
          divDom = document.createElement('div');
          divDom.classList.add('video-line');
          divDom.append(videoDom);
          document.getElementById('video-wrap').append(divDom);
          return;
        }

        videoLineDoms[1].append(videoDom);
        return;
      }
    },

    removeVideo: function (peerId) {
      const videoDom = document.getElementById(peerId);
      videoDom.remove();
    }
  }
};
</script>

<style lang="scss">
.video {
  width: 100vw;
  padding: 10px 0;
  display: flex;
  flex-wrap: wrap;
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
  }

  &-individual {
    flex: 1;
    margin: 0 10px;
    border-radius: 80px;

    &-focus {
      border: solid 5px orange;
    }
  }
}
</style>
