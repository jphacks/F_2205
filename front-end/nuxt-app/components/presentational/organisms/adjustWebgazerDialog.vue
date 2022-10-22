<template>
  <v-row justify="center">
    <v-dialog v-model="isOpenAdjustWebGazerDialog" fullscreen>
      <v-card class="cardContainer">
        <v-toolbar>
          <v-btn icon dark @click="handleAdjustWebGazer">
            <v-icon>mdi-close</v-icon>
          </v-btn>
          <v-toolbar-title>Settings</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-toolbar-items>
            <v-menu left bottom :close-on-content-click="true">
              <template v-slot:activator="{ on, attrs }">
                <v-btn icon v-bind="attrs" v-on="on">
                  <v-icon>mdi-dots-vertical</v-icon>
                </v-btn>
              </template>
              <v-list>
                <v-list-item>
                  <v-list-item-icon>
                    <v-icon>mdi-eye-check</v-icon>
                  </v-list-item-icon>
                  <v-list-item-content>
                    <v-list-item-title>現在の精度:{{ this.precisionMeasurement }}%</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item class="readjustButton" @click="readjustPosition">
                  <v-list-item-icon>
                    <v-icon>mdi-crosshairs</v-icon>
                  </v-list-item-icon>
                  <v-list-item-content>
                    <v-list-item-title> もう一度調整する </v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item class="readjustButton">
                  <v-list-item-icon>
                    <v-icon>mdi-close</v-icon>
                  </v-list-item-icon>
                  <v-list-item-content>
                    <v-list-item-title>閉じる</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
              </v-list>
            </v-menu>
          </v-toolbar-items>
        </v-toolbar>
        <v-row justify="center" align="center" class="adjustFieldContainer">
          <v-col cols="12" sm="8" md="10" class="adjustCanvus">
            <v-row justify="space-between" align="center" class="buttons">
              <v-col cols="12" sm="1"></v-col>
              <ClickAdjustBtn @add-adjustpoint="adjustPosition" :isExplainClickPoint="this.isExplainClickPointDialog" />
              <ClickAdjustBtn @add-adjustpoint="adjustPosition" :isExplainClickPoint="this.isExplainClickPointDialog" />
              <ClickAdjustBtn @add-adjustpoint="adjustPosition" :isExplainClickPoint="this.isExplainClickPointDialog" />
            </v-row>
            <v-row justify="space-between" align="center" class="buttons">
              <ClickAdjustBtn @add-adjustpoint="adjustPosition" :isExplainClickPoint="this.isExplainClickPointDialog" />
              <v-col cols="12" sm="1">
                <ClickAdjustBtn
                  @add-adjustpoint="adjustPosition"
                  v-if="adjustPoint > 8"
                  :isExplainClickPoint="this.isExplainClickPointDialog"
                />
                <v-col cols="12" sm="1" class="centerLabel">{{ storePredictionPoint }}{{ adjustPoint }}</v-col>
              </v-col>
              <ClickAdjustBtn @add-adjustpoint="adjustPosition" :isExplainClickPoint="this.isExplainClickPointDialog" />
            </v-row>
            <v-row justify="space-between" align="center" class="buttons">
              <ClickAdjustBtn @add-adjustpoint="adjustPosition" :isExplainClickPoint="this.isExplainClickPointDialog" />
              <ClickAdjustBtn @add-adjustpoint="adjustPosition" :isExplainClickPoint="this.isExplainClickPointDialog" />
              <ClickAdjustBtn @add-adjustpoint="adjustPosition" :isExplainClickPoint="this.isExplainClickPointDialog" />
              <ClickAdjustBtn @add-adjustpoint="adjustPosition" :isExplainClickPoint="this.isExplainClickPointDialog" />
            </v-row>
            <ExplainClickPointDialog
              :isOpenExplainClickPointDialog="this.isExplainClickPointDialog"
              @close-explain-click-pointDialog="closeExplainClickPointDialog"
            />
            <GazeCenterPointDialog
              :isOpenGazeCenterPointDialog="this.isGazeCenterPointDialog"
              @close-gazecenterpointdialog="closeGazeCenterPointDialog"
            />
            <LearningResultDialog
              :isOpenLearningResultDialog="this.isLearningResultDialog"
              :recisionMeasurement="this.precisionMeasurement"
              @close-learningresultdialog="closeLearningResultDialog"
            />
          </v-col>
        </v-row>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<style lang="scss">
.cardContainer {
  height: 100%;
}
.gazerContainer {
  height: 100%;
}
.adjustFieldContainer {
  height: 100%;
}
.buttons {
  height: 33%;
}
.adjustCanvus {
  height: 100%;
}
.centerLabel {
  padding-left: 27px;
}
.readjustButton {
  &:hover {
    cursor: pointer;
  }
}
</style>

<script>
import ClickAdjustBtn from '~/components/adjustWebgazer/atoms/clickAdjustBtn';
import ExplainClickPointDialog from '~/components/adjustWebgazer/organisms/explainClickPointDialog';
import GazeCenterPointDialog from '~/components/adjustWebgazer/organisms/gazeCenterPointDialog';
import LearningResultDialog from '~/components/adjustWebgazer/organisms/learningResultDialog';
import calculatePrecision from '~/components/adjustWebgazer/script/precision_calculation';

import webgazer from 'webgazer';
export default {
  props: ['isOpenAdjustWebGazerDialog', 'handleAdjustWebGazer'],
  components: {
    ClickAdjustBtn,
    ExplainClickPointDialog,
    GazeCenterPointDialog,
    LearningResultDialog
  },
  data() {
    return {
      xprediction: '',
      yprediction: '',
      adjustPoint: 0,
      isExplainClickPointDialog: true,
      isGazeCenterPointDialog: false,
      isLearningResultDialog: false,
      isStartStorePredictionPoint: false,
      screenWidth: window.innerWidth,
      screenHeight: window.innerHeight,
      precisionMeasurement: ''
    };
  },

  mounted: async function () {
    // webgazer.clearData();
    // webgazer.setRegression('ridge').setTracker('clmtrackr').begin();
    // webgazer.applyKalmanFilter(true).setGazeListener((data, clock) => {
    // });
    // var setup = function () {
    //   //Set up the main canvas. The main canvas is used to calibrate the webgazer.
    //   var canvas = document.getElementById('plotting_canvas');
    //   canvas.width = window.innerWidth;
    //   canvas.height = window.innerHeight;
    //   canvas.style.position = 'fixed';
    // };
    // setup();
    // webgazer.showVideoPreview(true).showPredictionPoints(true).applyKalmanFilter(true);
  },

  watch: {
    isOpenAdjustWebGazerDialog: function (newAdjust, oldAdjust) {
      if (newAdjust) {
        var setup = function () {
          //Set up the main canvas. The main canvas is used to calibrate the webgazer.
          var canvas = document.getElementById('plotting_canvas');
          console.log(canvas);
          canvas.width = window.innerWidth;
          canvas.height = window.innerHeight;
          canvas.style.position = 'fixed';
        };
        setup();

        this.isExplainClickPointDialog = true;
        this.adjustPoint = 0;
      }
    }
  },
  computed: {
    storePredictionPoint() {
      if (this.isStartStorePredictionPoint) {
        var canvas = document.getElementById('plotting_canvas');
        canvas.getContext('2d').clearRect(0, 0, canvas.width, canvas.height);

        console.log('storePredictionPoint');
        webgazer.params.storingPoints = true;

        this.sleep(5000).then(() => {
          webgazer.params.storingPoints = false;
          var pastData = webgazer.getStoredPoints();
          this.precisionMeasurement = calculatePrecision(pastData, this.screenHeight, this.screenWidth);

          var canvas = document.getElementById('plotting_canvas');
          canvas.getContext('2d').clearRect(0, 0, canvas.width, canvas.height);

          this.isLearningResultDialog = true;
          this.isStartStorePredictionPoint = false;
          return 'success';
        });
      }
      return 'center';
    }
  },
  methods: {
    stopWebgather() {
      webgazer.end();
      console.log('stop');
    },
    adjustPosition() {
      this.adjustPoint = this.adjustPoint + 1;
      if (this.adjustPoint === 10) {
        this.isGazeCenterPointDialog = true;
      }
    },
    sleep(time) {
      return new Promise((resolve, reject) => {
        setTimeout(() => {
          resolve();
        }, time);
      });
    },
    closeGazeCenterPointDialog() {
      this.isGazeCenterPointDialog = false;
      this.isStartStorePredictionPoint = true;
    },
    closeLearningResultDialog() {
      this.isLearningResultDialog = false;
    },
    closeExplainClickPointDialog() {
      this.isExplainClickPointDialog = false;
    },
    readjustPosition() {
      //   webgazer.pause();

      //   webgazer.clearData();
      this.isExplainClickPointDialog = true;
      this.adjustPoint = 0;

      //   webgazer.resume();
    }
  }
};
</script>
