<template>
  <section class="video-effect" :id="'video' + videoId + '-effect'">
    <section class="video-effect-waiwai" v-if="effectNumber == '1'">
      <img src="~/assets/img/waiwaiImg.svg" class="video-effect-waiwai-img-1" />
      <img src="~/assets/img/waiwaiImg.svg" class="video-effect-waiwai-img-2" />
      <img src="~/assets/img/waiwaiImg.svg" class="video-effect-waiwai-img-3" />
      <img src="~/assets/img/waiwaiImg.svg" class="video-effect-waiwai-img-4" />
    </section>

    <section class="video-effect-gogogo" v-if="effectNumber == '2'">
      <img src="~/assets/img/gogogoImg.svg" class="video-effect-gogogo-img-1" />
      <img src="~/assets/img/gogogoImg.svg" class="video-effect-gogogo-img-2" />
      <img src="~/assets/img/gogogoImg.svg" class="video-effect-gogogo-img-3" />
    </section>

    <section class="video-effect-goku" v-if="effectNumber == '3'">
      <img src="~/assets/img/gokuImg.svg" class="video-effect-goku-img-1" />
      <img src="~/assets/img/gokuImg.svg" class="video-effect-goku-img-2" />
      <img src="~/assets/img/gokuImg.svg" class="video-effect-goku-img-3" />
    </section>

    <section class="video-effect-drink-too-much" v-if="effectNumber == '4'">
      <div class="video-effect-drink-too-much-back"></div>
      <v-icon color="yellow" size="102" class="video-effect-drink-too-much-icon">mdi-alert </v-icon>
      <h3 class="video-effect-drink-too-much-ms">飲みすぎ注意！</h3>
    </section>

    <section class="video-effect-uoo" v-if="effectNumber == '5'">
      <img src="~/assets/img/shockImg.svg" class="video-effect-uoo-back" />
      <img src="~/assets/img/uooImg.svg" class="video-effect-uoo-img-1" />
      <img src="~/assets/img/uooImg.svg" class="video-effect-uoo-img-2" />
    </section>

    <section class="video-effect-iine" v-if="effectNumber == '6'">
      <img src="~/assets/img/onnpuImg.svg" class="video-effect-iine-back" />
      <img src="~/assets/img/onnpuImg.svg" class="video-effect-iine-back-2" />
      <img src="~/assets/img/iineImg.svg" class="video-effect-iine-img-1" />
    </section>

    <section class="video-effect-rest-room" v-if="isRestRoom">
      <img src="~/assets/img/flowerImg.jpg" class="video-effect-rest-room-img" />
      <h3 class="video-effect-rest-room-ms">～お花を摘んでいます～</h3>
    </section>
  </section>
</template>

<script>
export default {
  props: ['videoId'],
  data() {
    return {
      effectNumber: null,
      isCurrentEffect: false,
      isRestRoom: false
    };
  },
  methods: {
    start: function (effectNumber) {
      if (this.isCurrentEffect && this.effectNumber != 4) return;

      this.isCurrentEffect = true;
      this.effectNumber = effectNumber;

      const videoDom = document.querySelector(`#video${this.videoId}`);
      const width = videoDom.offsetWidth;
      const height = videoDom.offsetHeight;
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
    restRoomStart: function () {
      if (this.isCurrentEffect) return;

      this.effectNumber = null;
      this.isRestRoom = true;
      this.isCurrentEffect = true;

      const videoDom = document.querySelector(`#video${this.videoId}`);
      const width = videoDom.offsetWidth;
      const height = videoDom.offsetHeight;
      const x = videoDom.getBoundingClientRect().left;
      const y = videoDom.getBoundingClientRect().top;

      const tgDom = document.querySelector(`#video${this.videoId}-effect`);

      tgDom.classList.add('open');
      tgDom.style.width = width + 'px';
      tgDom.style.height = height + 'px';
      tgDom.style.top = y + 'px';
      tgDom.style.left = x + 'px';
    },
    restRoomEnd: function () {
      document.querySelector(`#video${this.videoId}-effect`).classList.remove('open');
      this.isRestRoom = false;
      this.isCurrentEffect = false;
    },
    resize: function () {
      if (!this.isCurrentEffect) return;

      const videoDom = document.querySelector(`#video${this.videoId}`);
      const width = videoDom.offsetWidth;
      const height = videoDom.offsetHeight;
      const x = videoDom.getBoundingClientRect().left;
      const y = videoDom.getBoundingClientRect().top;

      const tgDom = document.querySelector(`#video${this.videoId}-effect`);

      tgDom.style.width = width + 'px';
      tgDom.style.height = height + 'px';
      tgDom.style.top = y + 'px';
      tgDom.style.left = x + 'px';
    },
    remove: function () {
      const tgDom = document.querySelector(`#video${this.videoId}-effect`);
      tgDom.classList.remove('open');
      this.isCurrentEffect = false;
    },
    thisDomRemove: function () {
      const tgDom = document.querySelector(`#video${this.videoId}-effect`);
      tgDom.remove();
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
  position: fixed;
  top: 0;
  left: 0;
  width: 0;
  height: 0;
  border-radius: 20px;
  opacity: 0;
  transition: opacity 0.4s;
  z-index: -100;
  overflow: hidden;

  &-waiwai {
    &-img-1 {
      position: absolute;
      top: 0;
      left: 5%;
      width: 35%;
      transform: rotate(10deg);
      animation: waiwai linear 0.4s infinite;
    }
    &-img-2 {
      position: absolute;
      top: 5%;
      right: 5%;
      width: 40%;
      transform: rotate(-10deg);
      animation: waiwai linear 0.4s infinite;
    }
    &-img-3 {
      position: absolute;
      bottom: 0;
      left: 10%;
      width: 40%;
      animation: waiwai linear 0.4s infinite;
    }
    &-img-4 {
      position: absolute;
      bottom: 5%;
      right: 5%;
      width: 30%;
      animation: waiwai linear 0.4s infinite;
    }
  }

  &-gogogo {
    &-img-1 {
      position: absolute;
      top: 0;
      left: 5%;
      width: 40%;
      transform: rotate(10deg);
      animation: gogogo linear 0.4s infinite;
    }
    &-img-2 {
      position: absolute;
      top: 10%;
      right: 0;
      width: 30%;
      transform: rotate(10deg);
      animation: gogogo linear 0.4s infinite;
    }
    &-img-3 {
      position: absolute;
      top: 50%;
      left: 10%;
      width: 30%;
      transform: rotate(10deg);
      animation: gogogo linear 0.4s infinite;
    }
  }

  &-goku {
    &-img-1 {
      position: absolute;
      top: 30%;
      left: 5%;
      width: 40%;
      transform: rotate(10deg);
      animation: goku linear 0.4s infinite;
    }
    &-img-2 {
      position: absolute;
      top: 40%;
      right: 5%;
      width: 40%;
      transform: rotate(10deg);
      animation: goku linear 0.4s infinite;
    }
    &-img-3 {
      position: absolute;
      top: 0%;
      left: 30%;
      width: 30%;
      transform: rotate(10deg);
      animation: goku linear 0.4s infinite;
    }
  }

  &-drink-too-much {
    width: 100%;
    height: 100%;
    &-icon {
      padding: 20% 0 5%;
      width: 100%;
      position: absolute;
      animation: drink-too-much-icon linear 1s infinite;
    }

    &-ms {
      width: 100%;
      text-align: center;
      color: yellow;
    }

    &-back {
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background-color: orange;
      opacity: 0.2;
    }
  }

  &-uoo {
    &-back {
      width: 100%;
    }

    &-img-1 {
      width: 40%;
      position: absolute;
      top: 0%;
      left: 6%;
      transform: rotate(-20deg);
      animation: uoo-img-1 linear 0.4s infinite;
    }
    &-img-2 {
      width: 30%;
      position: absolute;
      top: 7%;
      right: 6%;
      transform: rotate(10deg);
      animation: uoo-img-2 linear 0.4s infinite;
    }
  }

  &-iine {
    &-back {
      width: 20%;
      position: absolute;
      top: 15%;
      right: 25%;
      transform: rotate(20deg);
      animation: iine-back linear 0.4s infinite;
    }
    &-back-2 {
      width: 15%;
      position: absolute;
      top: 5%;
      right: 5%;
      transform: rotate(10deg);
      animation: iine-back-2 linear 0.4s infinite;
    }

    &-img-1 {
      width: 42%;
      position: absolute;
      top: 0%;
      left: 6%;
      transform: rotate(-20deg);
      animation: iine linear 0.4s infinite;
    }
  }

  &-rest-room {
    &-img {
      width: 130%;
      opacity: 0.7;
    }
    &-ms {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      font-size: 18px;
      text-align: center;
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

@keyframes gogogo {
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

@keyframes goku {
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

@keyframes drink-too-much-icon {
  0% {
    transform: scale(1);
  }

  25% {
    transform: scale(1.1);
  }

  50% {
    transform: scale(1);
  }

  75% {
    transform: scale(0.9);
  }
  100% {
    transform: scale(1);
  }
}

@keyframes uoo-img-1 {
  0% {
    transform: translateY(0) rotate(-20deg);
  }

  50% {
    transform: translateY(7px) rotate(-20deg);
  }

  100% {
    transform: translateY(0) rotate(-20deg);
  }
}

@keyframes uoo-img-2 {
  0% {
    transform: translateY(0) rotate(20deg);
  }

  50% {
    transform: translateY(7px) rotate(20deg);
  }

  100% {
    transform: translateY(0) rotate(20deg);
  }
}

@keyframes iine {
  0% {
    transform: translateY(0) rotate(-20deg);
  }

  50% {
    transform: translateY(7px) rotate(-20deg);
  }

  100% {
    transform: translateY(0) rotate(-20deg);
  }
}

@keyframes iine-back {
  0% {
    transform: translateY(0) rotate(20deg);
  }

  50% {
    transform: translateY(7px) rotate(20deg);
  }

  100% {
    transform: translateY(0) rotate(20deg);
  }
}

@keyframes iine-back-2 {
  0% {
    transform: translateY(0) rotate(10deg);
  }

  50% {
    transform: translateY(7px) rotate(10deg);
  }

  100% {
    transform: translateY(0) rotate(10deg);
  }
}
</style>
