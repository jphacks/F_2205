<template>
  <v-card class="video-state" color="orange lighten-3" tile>
    <div class="video-state-inner pa-2 d-flex align-center">
      <div class="video-state-icon">
        <v-avatar color="white" size="56" class="mx-4 video-state-icon-avatar">
          <v-icon color="black" @click.native="focusThisVideoAllLiftFn">mdi-account-group </v-icon>
        </v-avatar>

        <v-avatar color="white" size="56" class="mx-4 video-state-icon-avatar">
          <v-icon color="black">mdi-microphone</v-icon>
        </v-avatar>

        <v-avatar color="white" size="56" class="mx-4 video-state-icon-avatar">
          <v-icon color="black">mdi-camera</v-icon>
        </v-avatar>

        <v-avatar color="white" size="56" class="mx-4 video-state-icon-avatar">
          <v-btn color="transparent" height="56" @click.native="gazeEstimatingFn">
            <v-icon color="black" v-if="isEnableGazeEstimating">mdi-eye-outline</v-icon>
            <v-icon color="black" v-else>mdi-eye-off-outline</v-icon>
          </v-btn>
        </v-avatar>
        <v-avatar color="white" size="56" class="mx-4 video-state-icon-avatar">
          <v-btn color="transparent" height="56" @click.native="handleAdjustWebGazer">
            <v-icon color="black">mdi-android</v-icon>
          </v-btn>
        </v-avatar>
      </div>

      <div>
        <AdjustWebgazerDialog
          :isOpenAdjustWebGazerDialog="this.isOpenAdjustWebGazerDialog"
          :handleAdjustWebGazer="this.handleAdjustWebGazer"
          v-if="isEnableGazeEstimating"
        />
      </div>

      <div>
        <Btn text="退出" color="red" :clickedfn="leavingFn" />
      </div>
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
    'isOpenAdjustWebGazerDialog'
  ],
  components: {
    Btn,
    AdjustWebgazerDialog
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
    }
  }
}
</style>
