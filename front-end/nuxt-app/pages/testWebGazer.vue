<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="10" class="adjustCanvus">
      <v-row justify="space-between" align="center" class="buttons">
        <ClickAdjustBtn @add-adjustpoint="adjustPosition" />
        <ClickAdjustBtn @add-adjustpoint="adjustPosition" />
        <ClickAdjustBtn @add-adjustpoint="adjustPosition" />
        <ClickAdjustBtn @add-adjustpoint="adjustPosition" />
      </v-row>
      <v-row justify="space-between" align="center" class="buttons">
        <ClickAdjustBtn @add-adjustpoint="adjustPosition" />
        <ClickAdjustBtn @add-adjustpoint="adjustPosition" v-if="adjustPoint > 2" />
        <ClickAdjustBtn @add-adjustpoint="adjustPosition" />
      </v-row>
      <v-row justify="space-between" align="center" class="buttons">
        <ClickAdjustBtn @add-adjustpoint="adjustPosition" />
        <ClickAdjustBtn @add-adjustpoint="adjustPosition" />
        <ClickAdjustBtn @add-adjustpoint="adjustPosition" />
        <ClickAdjustBtn @add-adjustpoint="adjustPosition" />
      </v-row>
      <!-- <v-card>
        <v-card-title class="headline"> Welcome to the Vuetify + Nuxt.js template </v-card-title>
        <div>x: {{ xprediction }}</div>
        <div>y: {{ yprediction }}</div>
        <div>tracker: {{ currentTracker }}</div>
        <div>regression: {{ currentRegression }}</div>
        <v-card-actions>
          <v-spacer />
          <v-btn color="primary" nuxt to="/inspire" @click="stopWebgather"> Continue </v-btn>
          <v-btn color="primary" @click="stopWebgather">stopWebgazer</v-btn>
        </v-card-actions>
      </v-card> -->
      <ExplainClickPointDialog />
      <GazeCenterPointDialog :isOpenDialog="this.isGazeCenterPointDialog" @close-dialog="closeGazeCenterPointDialog" />
    </v-col>
  </v-row>
</template>

<style lang="scss">
.buttons {
  height: 33%;
}
.adjustCanvus {
  height: 80%;
}
</style>

<script>
import webgazer from 'webgazer';
import ClickAdjustBtn from '~/components/adjustWebgazer/atoms/clickAdjustBtn';
import ExplainClickPointDialog from '~/components/adjustWebgazer/organisms/explainClickPointDialog';
import GazeCenterPointDialog from '~/components/adjustWebgazer/organisms/gazeCenterPointDialog';

export default {
  name: 'IndexPage',
  layout: 'testWebGazer',
  components: {
    ClickAdjustBtn,
    ExplainClickPointDialog,
    GazeCenterPointDialog
  },
  data() {
    return {
      cart: [],
      premium: false,
      hello: 'hello',
      webgatherData: '',
      webgatherClock: '',
      xprediction: '',
      yprediction: '',
      tracker: '',
      regression: '',
      adjustPoint: 0,
      isGazeCenterPointDialog: false
    };
  },

  mounted: async function () {
    // webgazer.setRegression('ridge').setTracker('clmtrackr').begin();
    // webgazer.applyKalmanFilter(true).setGazeListener((data, clock) => {
    //   this.xprediction = data.x;
    //   this.yprediction = data.y;
    // });
  },
  computed: {
    currentTracker() {
      return (this.tracker = webgazer.getTracker().name);
    },
    currentRegression() {
      return (this.regression = webgazer.getRegression()[0].name);
    }
  },
  methods: {
    stopWebgather() {
      webgazer.end();
      console.log('stop');
    },
    adjustPosition() {
      this.adjustPoint = this.adjustPoint + 1;
      if (this.adjustPoint === 3) {
        this.isGazeCenterPointDialog = true;
      }
    },
    closeGazeCenterPointDialog() {
      this.isGazeCenterPointDialog = false;
    }
  }
};
</script>
