<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="10" class="adjustCanvus">
      <v-row justify="space-between" align="center" class="buttons">
        <ClickAdjustBtn />
        <ClickAdjustBtn />
        <ClickAdjustBtn />
        <ClickAdjustBtn />
      </v-row>
      <v-row justify="space-between" align="center" class="buttons">
        <ClickAdjustBtn />
        <ClickAdjustBtn />
        <ClickAdjustBtn />
      </v-row>
      <v-row justify="space-between" align="center" class="buttons">
        <ClickAdjustBtn />
        <ClickAdjustBtn />
        <ClickAdjustBtn />
        <ClickAdjustBtn />
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

export default {
  name: 'IndexPage',
  layout: 'testWebGazer',
  components: {
    ClickAdjustBtn
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
      regression: ''
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
    }
  }
};
</script>
