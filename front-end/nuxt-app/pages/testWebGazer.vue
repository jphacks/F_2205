<template>
  <div class="gazerContainer">
    <v-app-bar color="deep-purple accent-4" dense dark>
      <v-app-bar-nav-icon></v-app-bar-nav-icon>
      <v-toolbar-title>NomiPara</v-toolbar-title>
      <v-spacer></v-spacer>
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
    </v-app-bar>
    <div id="capture">
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
          <div>
            <v-btn id="save-btn" @click="captureImage"> 画像を保存 </v-btn>
          </div>
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
    </div>
  </div>
</template>

<style lang="scss">
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
import webgazer from 'webgazer';
import html2canvas from 'html2canvas';

import ClickAdjustBtn from '~/components/presentational/atoms/adjustWebgazer/clickAdjustBtn';
import ExplainClickPointDialog from '~/components/presentational/organisms/adjustWebgazer/explainClickPointDialog';
import GazeCenterPointDialog from '~/components/presentational/organisms/adjustWebgazer/gazeCenterPointDialog';
import LearningResultDialog from '~/components/presentational/organisms/adjustWebgazer/learningResultDialog';
import calculatePrecision from '../script/precision_calculation';

export default {
  name: 'IndexPage',
  layout: 'testWebGazer',
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
    webgazer.clearData();
    webgazer.setRegression('ridge').setTracker('clmtrackr').begin();
    webgazer.applyKalmanFilter(true).setGazeListener((data, clock) => {
      // this.xprediction = data.x;
      // this.yprediction = data.y;
    });

    var setup = function () {
      //Set up the main canvas. The main canvas is used to calibrate the webgazer.
      var canvas = document.getElementById('plotting_canvas');
      canvas.width = window.innerWidth;
      canvas.height = window.innerHeight;
      canvas.style.position = 'fixed';
    };
    setup();

    // webgazer.showVideoPreview(true).showPredictionPoints(true).applyKalmanFilter(true);
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

          console.log(pastData, this.precisionMeasurement);

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
    captureImage() {
      html2canvas(document.querySelector('#capture')).then((canvas) => {
        const link = document.createElement('a');
        link.href = canvas.toDataURL();
        link.download = `export_image.png`;
        link.click();
      });
    },
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
      webgazer.pause();

      webgazer.clearData();
      this.isExplainClickPointDialog = true;
      this.adjustPoint = 0;

      webgazer.resume();
    }
  }
};
</script>
