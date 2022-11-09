<template>
  <section class="video-name open" :id="'video' + videoId + '-name'">
    {{ videoName }}
  </section>
</template>

<script>
export default {
  props: ['videoId'],
  data() {
    return {
      videoName: null
    };
  },
  methods: {
    setVideoName: function () {
      if (this.videoId == 'my-video') {
        this.videoName = 'あなた';
        this.setPosition();
        return;
      }

      if (localStorage.getItem(`video${this.videoId}`) != null) {
        this.videoName = localStorage.getItem(`video${this.videoId}`);
      } else {
        this.videoName = 'unknown';
      }

      this.setPosition();
    },
    setPositionWaitDisplay() {
      // 位置調整が終わるまで非表示にする
      document.querySelector(`#video${this.videoId}-name`).classList.remove('open');
    },
    setPosition: function () {
      const videoDom = document.querySelector(`#video${this.videoId}`);
      const x = videoDom.getBoundingClientRect().left;
      const y = videoDom.getBoundingClientRect().top;

      const tgDom = document.querySelector(`#video${this.videoId}-name`);
      tgDom.style.top = y + 'px';
      tgDom.style.left = x + 'px';

      document.querySelector(`#video${this.videoId}-name`).classList.add('open');
    },
    remove: function () {
      const tgDom = document.querySelector(`#video${this.videoId}-name`);
      tgDom.remove();
    }
  },
  mounted: function () {
    // websocketからmemberデータが送られてくるタイミングを考慮して時間差を作る
    // skyweay → websocket の順番を担保する 万が一逆になった場合はunknown
    // 既存ビデオの位置がずれる時間の考慮も含む
    this.setPositionWaitDisplay();
    setTimeout(this.setVideoName, 5000);
  }
};
</script>

<style lang="scss">
.open.video-name {
  opacity: 0.8;
}
.video-name {
  position: fixed;
  top: 0;
  left: 0;
  font-size: 14px;
  padding: 2px 15px;
  border-radius: 20px 0 20px 0;
  background-color: #ff8f00;
  opacity: 0;
}
</style>
