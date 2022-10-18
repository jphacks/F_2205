<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card class="logo py-4 d-flex justify-center">
        <NuxtLogo />
        <VuetifyLogo />
      </v-card>
      <v-card>
        <v-card-title class="headline"> Welcome to the Vuetify + Nuxt.js template </v-card-title>
        <div>x: {{ xprediction }}</div>
        <div>y: {{ yprediction }}</div>
        <div>tracker: {{ currentTracker }}</div>
        <div>regression: {{ currentRegression }}</div>
        <v-card-actions>
          <v-spacer />
          <v-btn color="primary" nuxt to="/inspire"> Continue </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import webgazer from 'webgazer';

export default {
  name: 'IndexPage',
  layout: 'testWebGazer',
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
    webgazer.setRegression('ridge').setTracker('clmtrackr').begin();
    webgazer.applyKalmanFilter(true).setGazeListener((data, clock) => {
      this.xprediction = data.x;
      this.yprediction = data.y;
    });
  },

  computed: {
    currentTracker() {
      return (this.tracker = webgazer.getTracker().name);
    },
    currentRegression() {
      return (this.regression = webgazer.getRegression()[0].name);
    }
  }
};
</script>
