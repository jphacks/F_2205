<template>
  <section id="video-wrap" class="video">
    <div class="video-line">
      <div class="video-individual">
        <video id="my-video" autoplay muted playsinline></video>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  props: ['roomMemberNum'],

  methods: {
    addVideo: function (stream, roomMemberNum) {
      const videoLineDoms = document.querySelectorAll('.video-line');

      const videoContainer = document.createElement('div');
      videoContainer.classList.add('video-individual');

      const videoDom = document.createElement('video');
      videoDom.setAttribute('id', stream.peerId);
      // videoDom.classList.add('video-individual');
      videoDom.srcObject = stream;
      videoDom.play();
      videoContainer.appendChild(videoDom);

      //append
      let divDom;
      if (roomMemberNum <= 3) {
        videoLineDoms[0].append(videoContainer);
        return;
      }

      if (roomMemberNum > 3 && roomMemberNum <= 6) {
        if (roomMemberNum == 4) {
          divDom = document.createElement('div');
          divDom.classList.add('video-line');
          divDom.append(videoContainer);
          document.getElementById('video-wrap').append(divDom);
          return;
        }

        videoLineDoms[1].append(videoContainer);
        return;
      }

      if (roomMemberNum > 6 && roomMemberNum <= 9) {
        if (roomMemberNum == 7) {
          divDom = document.createElement('div');
          divDom.classList.add('video-line');
          divDom.append(videoContainer);
          document.getElementById('video-wrap').append(divDom);
          return;
        }

        videoLineDoms[2].append(videoContainer);
        return;
      }

      if (roomMemberNum > 9 && roomMemberNum <= 12) {
        if (roomMemberNum == 10) {
          divDom = document.createElement('div');
          divDom.classList.add('video-line');
          divDom.append(videoContainer);
          document.getElementById('video-wrap').append(divDom);
          return;
        }

        videoLineDoms[3].append(videoContainer);
        return;
      }

      if (roomMemberNum > 12 && roomMemberNum <= 15) {
        if (roomMemberNum == 13) {
          divDom = document.createElement('div');
          divDom.classList.add('video-line');
          divDom.append(videoContainer);
          document.getElementById('video-wrap').append(divDom);
          return;
        }

        videoLineDoms[4].append(videoContainer);
        return;
      }

      if (roomMemberNum > 15 && roomMemberNum <= 18) {
        if (roomMemberNum == 16) {
          divDom = document.createElement('div');
          divDom.classList.add('video-line');
          divDom.append(videoContainer);
          document.getElementById('video-wrap').append(divDom);
          return;
        }

        videoLineDoms[5].append(videoContainer);
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
    * {
      flex: 1;
      margin: 10px;
      border-radius: 20px;
      max-width: 480px;
    }

    &-focus {
      border: solid 5px orange;
    }
  }
}
</style>
