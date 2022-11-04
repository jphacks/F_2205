<template>
  <v-card class="video-state" color="orange lighten-3" tile>
    <div class="video-state-inner pa-2 d-flex align-center">
      <div class="video-state-icon">
        <v-avatar color="white" size="56" class="mx-4 video-state-icon-avatar effect-btn">
          <v-icon color="black" @click="this.effectOperation" id="effect-icon">mdi-thumb-up </v-icon>
        </v-avatar>

        <v-avatar color="white" size="56" class="mx-4 video-state-icon-avatar">
          <v-icon color="black" @click.native="focusThisVideoAllLiftFn">mdi-account-group</v-icon>
        </v-avatar>

        <v-avatar color="white" size="56" class="mx-4 video-state-icon-avatar">
          <v-icon color="black" @click.native="audioMuteFn" v-if="myAudioStatus">mdi-microphone</v-icon>
          <v-icon color="red" @click.native="audioMuteFn" v-if="!myAudioStatus">mdi-microphone-off</v-icon>
        </v-avatar>

        <v-avatar color="white" size="56" class="mx-4 video-state-icon-avatar">
          <v-icon color="black" @click.native="videoMuteFn" v-if="myVideoStatus">mdi-camera</v-icon>
          <v-icon color="red" @click.native="videoMuteFn" v-if="!myVideoStatus">mdi-camera-off</v-icon>
        </v-avatar>

        <v-avatar color="white" size="56" class="mx-4 video-state-icon-avatar">
          <v-btn color="transparent" height="56" @click.native="gazeEstimatingFn">
            <v-icon color="black" v-if="isEnableGazeEstimating">mdi-eye-outline</v-icon>
            <v-icon color="black" v-else>mdi-eye-off-outline</v-icon>
          </v-btn>
        </v-avatar>
        <v-avatar color="white" size="56" class="mx-4 video-state-icon-avatar">
          <v-btn color="transparent" height="56" @click.native="handleAdjustWebGazer" v-if="isEnableGazeEstimating">
            <v-icon color="black">mdi-cog-outline</v-icon>
          </v-btn>
          <v-btn color="transparent" height="56" v-else>
            <v-icon color="black">mdi-cog-off-outline</v-icon>
          </v-btn>
        </v-avatar>
        <v-avatar color="white" size="56" class="mx-4 video-state-icon-avatar">
          <v-icon color="black" @click.native="captureImage">mdi-monitor-screenshot</v-icon>
        </v-avatar>
      </div>

      <div class="adjustWebgazerContainer">
        <AdjustWebgazerDialog
          :isOpenAdjustWebGazerDialog="this.isOpenAdjustWebGazerDialog"
          :handleAdjustWebGazer="this.handleAdjustWebGazer"
          v-if="isOpenAdjustWebGazerDialog"
          class="adjustWebgazer"
        />
      </div>

      <div>
        <Btn text="退出" color="red" :clickedfn="leavingFn" />
      </div>
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
  &-inner {
    max-width: 1240px;
    margin: 0 auto;
  }
  &-icon {
    text-align: center;
    width: 100%;
    transform: translateX(5%);

    &-avatar {
      box-shadow: 2px 4px 6px 0 rgba(0, 0, 0, 0.2);
      &:hover {
        cursor: pointer;
        opacity: 0.9;
      }
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
  opacity: 1;
}
</style>
