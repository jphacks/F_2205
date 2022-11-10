<template>
  <v-card class="video-state" tile>
    <div class="video-state-inner pa-2 d-flex align-center">
      <div class="video-state-icon">
        <div class="video-state-icon-wrap">
          <v-icon color="white" size="42" @click.native="captureImage">mdi-monitor-screenshot</v-icon>
          <h3 class="video-state-icon-info">撮影</h3>
        </div>

        <div class="video-state-icon-wrap">
          <v-icon color="white" size="42" @click="this.effectOperation" id="effect-icon">mdi-thumb-up</v-icon>
          <h3 class="video-state-icon-info">エフェクト</h3>
        </div>

        <div class="video-state-icon-wrap">
          <v-icon color="white" size="42" style="margin-right: -10px !important">mdi-human-female</v-icon>
          <v-icon color="white" size="42" style="margin-left: -10px !important">mdi-human-male</v-icon>
          <h3 class="video-state-icon-info">トイレ</h3>
        </div>

        <div class="video-state-icon-wrap">
          <v-icon color="white" size="42" @click.native="audioMuteFn" v-if="myAudioStatus">mdi-microphone</v-icon>
          <v-icon color="red" size="42" @click.native="audioMuteFn" v-if="!myAudioStatus">mdi-microphone-off</v-icon>
          <h3 class="video-state-icon-info">マイク</h3>
        </div>

        <div class="video-state-icon-wrap">
          <v-icon color="white" size="42" @click.native="videoMuteFn" v-if="myVideoStatus">mdi-camera</v-icon>
          <v-icon color="red" size="42" @click.native="videoMuteFn" v-if="!myVideoStatus">mdi-camera-off</v-icon>
          <h3 class="video-state-icon-info">カメラ</h3>
        </div>

        <div class="video-state-icon-wrap">
          <v-icon color="white" size="42" v-if="isEnableGazeEstimating" @click.native="gazeEstimatingFn"
            >mdi-eye-outline</v-icon
          >
          <v-icon color="red" size="42" v-else @click.native="gazeEstimatingFn">mdi-eye-off-outline</v-icon>
          <h3 class="video-state-icon-info">視線制御</h3>
        </div>

        <div class="video-state-icon-wrap">
          <v-icon color="white" size="42" @click.native="handleAdjustWebGazer" v-if="isEnableGazeEstimating"
            >mdi-cog-outline</v-icon
          >
          <v-icon color="red" size="42" v-else @click.native="handleAdjustWebGazer">mdi-cog-off-outline</v-icon>
          <h3 class="video-state-icon-info">視線調整</h3>
        </div>

        <div class="video-state-icon-wrap">
          <v-icon color="white" size="42" v-if="isEnableDrinkEstimating" @click.native="drinkEstimatingFn"
            >mdi-glass-cocktail</v-icon
          >
          <v-icon color="red" size="42" v-else @click.native="drinkEstimatingFn">mdi-glass-cocktail-off</v-icon>
          <h3 class="video-state-icon-info">飲み検知</h3>
        </div>
      </div>

      <div class="video-state-icon-wrap">
        <v-icon color="red" size="42" @click.native="leavingFn">mdi-exit-run </v-icon>
        <h3 class="video-state-icon-info">退出</h3>
      </div>

      <div class="adjustWebgazerContainer">
        <AdjustWebgazerDialog
          :isOpenAdjustWebGazerDialog="this.isOpenAdjustWebGazerDialog"
          :handleAdjustWebGazer="this.handleAdjustWebGazer"
          v-if="isOpenAdjustWebGazerDialog"
          class="adjustWebgazer"
        />
      </div>

      <!-- <div>
        <Btn text="退出" color="red" :clickedfn="leavingFn" />
      </div> -->
    </div>

    <div class="effect-select-box">
      <img src="~/assets/img/waiwaiImg.svg" class="effect-select-box-img" @click="effectFn('1')" />
      <img src="~/assets/img/gogogoImg.svg" class="effect-select-box-img" @click="effectFn('2')" />
      <img src="~/assets/img/gokuImg.svg" class="effect-select-box-img" @click="effectFn('3')" />
    </div>
  </v-card>
</template>

<script>
import Btn from '~/components/presentational/atoms/btn';
import AdjustWebgazerDialog from '~/components/presentational/organisms/adjustWebgazerDialog';

export default {
  props: [
    'leavingFn',
    'isEnableGazeEstimating',
    'gazeEstimatingFn',
    'isEnableDrinkEstimating',
    'drinkEstimatingFn',
    'focusThisVideoAllLiftFn',
    'handleAdjustWebGazer',
    'isOpenAdjustWebGazerDialog',
    'videoMuteFn',
    'audioMuteFn',
    'myAudioStatus',
    'myVideoStatus',
    'captureImage',
    'effectFn'
  ],
  components: {
    Btn,
    AdjustWebgazerDialog
  },

  methods: {
    effectOperation: function () {
      this.isEffectDisplay = !this.isEffectDisplay;

      const effectSelectBoxDom = document.querySelector('.effect-select-box');
      effectSelectBoxDom.classList.toggle('effect-open');

      const tgIconDom = document.querySelector('#effect-icon');
      const x = tgIconDom.getBoundingClientRect().left;
      const y = tgIconDom.getBoundingClientRect().top;

      effectSelectBoxDom.style.top = y - 100 + 'px';
      effectSelectBoxDom.style.left = x - 30 + 'px';
    }
  }
};
</script>

<style lang="scss">
.video-state {
  background-color: #ffc063 !important;
  &-inner {
    max-width: 1240px;
    margin: 0 auto;
  }
  &-icon {
    text-align: center;
    width: 100%;
    transform: translateX(5%);
    display: flex;
    justify-content: center;

    &-wrap {
      width: 115px;
      transition: opacity 0.2s;
      &:hover {
        cursor: pointer;
        opacity: 0.6;
      }
    }

    &-info {
      font-size: 14px;
    }
  }
}
.adjustWebgazerContainer {
  position: fixed;
  top: 0;
  left: 0;
}
.adjustWebgazer {
  width: 100vw;
  height: 100vh;
}

.effect-select-box {
  visibility: hidden;
  display: flex;
  align-items: center;
  justify-content: space-around;
  opacity: 0;
  position: fixed;
  top: 0;
  left: 0;
  width: 400px;
  height: 80px;
  border-radius: 20px !important;
  padding: 0 10px;
  background-color: #fff;
  transition: opacity 0.3s;

  &-img {
    width: 80px;
    &:hover {
      cursor: pointer;
      opacity: 0.6;
    }
  }
}

.effect-open.effect-select-box {
  visibility: visible;
  opacity: 1;
}
</style>
