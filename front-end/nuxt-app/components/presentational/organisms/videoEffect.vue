<template>
  <section class="video-effect" :id="'video' + videoId + '-effect'">
    <section class="video-effect-waiwai" v-if="effectNumber == '1'">
      <img src="~/assets/img/waiwaiImg.svg" class="video-effect-waiwai-img-1" />
      <img src="~/assets/img/waiwaiImg.svg" class="video-effect-waiwai-img-2" />
      <img src="~/assets/img/waiwaiImg.svg" class="video-effect-waiwai-img-3" />
    </section>
  </section>
</template>

<script>
export default {
  props: ['videoId'],
  data() {
    return {
      effectNumber: null
    };
  },
  methods: {
    start: function (effectNumber) {
      this.effectNumber = effectNumber;

      const videoDom = document.querySelector(`#video${this.videoId}`);
      const width = videoDom.clientWidth;
      const height = videoDom.clientHeight;
      const x = videoDom.getBoundingClientRect().left;
      const y = videoDom.getBoundingClientRect().top;

      const tgDom = document.querySelector(`#video${this.videoId}-effect`);

      tgDom.classList.add('open');
      tgDom.style.width = width + 'px';
      tgDom.style.height = height + 'px';
      tgDom.style.top = y + 'px';
      tgDom.style.left = x + 'px';

      setTimeout(this.remove, 5000);
    },
    remove: function () {
      const tgDom = document.querySelector(`#video${this.videoId}-effect`);
      tgDom.classList.remove('open');
    }
  },

  mounted: function () {}
};
</script>

<style lang="scss">
.open.video-effect {
  opacity: 1;
  z-index: 100;
}
.video-effect {
  //   background-color: red;
  position: fixed;
  top: 0;
  left: 0;
  width: 0;
  height: 0;
  border-radius: 20px;
  opacity: 0;
  transition: opacity 0.4s;
  z-index: -100;

  &-waiwai {
    &-img-1 {
      position: absolute;
      top: 0;
      left: 30px;
      width: 160px;
      transform: rotate(10deg);
      animation: waiwai linear 0.4s infinite;
    }
    &-img-2 {
      position: absolute;
      top: 80px;
      right: 30px;
      width: 160px;
      transform: rotate(-10deg);
      animation: waiwai linear 0.4s infinite;
    }
    &-img-3 {
      position: absolute;
      bottom: 10px;
      left: 60px;
      width: 160px;
      animation: waiwai linear 0.4s infinite;
    }
  }
}

@keyframes waiwai {
  0% {
    transform: translateY(0) rotate(10deg);
  }

  50% {
    transform: translateY(7px) rotate(9deg);
  }

  100% {
    transform: translateY(0) rotate(10deg);
  }
}
</style>
