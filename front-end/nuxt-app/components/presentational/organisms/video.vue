<template>
  <section id="video-wrap" class="video">
    <video id="my-video" class="video-individual" autoplay muted playsinline></video>
  </section>
</template>

<script>
export default {
  props: ['roomMemberNum'],

  methods: {
    addVideo: function (stream, roomMemberNum) {
      const divDom = document.createElement('div');
      divDom.classList.add('video-line');

      const videoDom = document.createElement('video');
      videoDom.setAttribute('id', stream.peerId);
      videoDom.classList.add('video-individual');
      videoDom.srcObject = stream;
      videoDom.play();
      document.getElementById('video-wrap').append(videoDom);

      console.log(roomMemberNum);

      //ビデオのリサイズ
      this.videoResize(roomMemberNum);
    },

    removeVideo: function (peerId, roomMemberNum) {
      const videoDom = document.getElementById(peerId);
      videoDom.remove();

      //ビデオのリサイズ
      this.videoResize(roomMemberNum);
    },

    videoResize: function (roomMemberNum) {
      //ビデオのリサイズ
      const videos = document.querySelectorAll('.video-individual');

      for (let video of videos) {
        switch (roomMemberNum) {
          case 1:
            video.style.width = '100%';
            video.style.height = '100%';
            break;
          case 2:
            video.style.width = '34%';
            video.style.height = 'auto';
            break;
          default:
            video.style.width = '34%';
            video.style.height = 'auto';
            break;
        }
      }
    }
  }
};
</script>

<style lang="scss">
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

  &-line {
    display: flex;
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
</style>
