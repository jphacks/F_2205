<template>
  <div>
    <div class=drinking>飲んだ秒数: {{ this.drinkingCount }}秒</div>
    <button v-on:click="initializeDrinkingModel()">Start</button>
    <button v-on:click="clearDrinkingCount()">Reset</button>
    <div id="webcam-container"></div>
    <div id="label-container"></div>
  </div>
</template>


<script>
import * as tf from '@tensorflow/tfjs';
import * as tmImage from '@teachablemachine/image';

export default {
  data() {
    return {
      model: null,
      webcam: null,
      labelContainer: null,
      maxPredictions: null,
      baseURL: 'https://teachablemachine.withgoogle.com/models/x8mEWPGTZ/',
      drinkingCount: 0, // 飲んだ回数
      predictionCount: 0, // 推定結果の返却回数
    }
  },
  methods: {
    initializeDrinkingModel: async function () {
      const modelURL = this.baseURL + "model.json";
      const metadataURL = this.baseURL + "metadata.json";

      // modelとmetadataのロード
      this.model = await tmImage.load(modelURL, metadataURL);
      this.maxPredictions = this.model.getTotalClasses();

      // webcamのセットアップ
      const flip = true;
      this.webcam = new tmImage.Webcam(500, 500, flip);
      await this.webcam.setup();
      await this.webcam.play();
      window.requestAnimationFrame(this.loop);

      // webcamで取得した情報をcanvasに表示
      document.getElementById("webcam-container").appendChild(this.webcam.canvas);
      this.labelContainer = document.getElementById("label-container");
      for (let i = 0; i < this.maxPredictions; i++) {
        this.labelContainer.appendChild(document.createElement("div"));
      }
    },
    loop: async function () {
      this.webcam.update();
      await this.predictDrinking();
      window.requestAnimationFrame(this.loop);
    },
    predictDrinking: async function () {
      // 推定結果
      const prediction = await this.model.predict(this.webcam.canvas);
      // 推定結果のテキスト表示
      for (let i = 0; i < this.maxPredictions; i++) {
        const classPrediction = prediction[i].className + ": " + prediction[i].probability.toFixed(2);
        this.labelContainer.childNodes[i].innerHTML = classPrediction;

        // Drinking class の精度が8割以上の時、カウントを行う
        if (prediction[1].probability.toFixed(2)>=0.80) {
          this.predictionCount += 1

          if (this.predictionCount > 100) {
            this.drinkingCount += 1;
            this.predictionCount = 0;
          } 
        }
      }
    },
    clearDrinkingCount: function () {
      this.drinkingCount = 0;
    }
  }
}
</script>

<style lang="scss">
</style>