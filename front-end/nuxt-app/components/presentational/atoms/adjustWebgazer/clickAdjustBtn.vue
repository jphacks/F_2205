<template>
  <v-col cols="12" sm="1">
    <v-btn class="mx-2" small fab dark :color="buttonLightenString" @click="subButtonLiten">
      <v-icon dark> mdi-android </v-icon>
    </v-btn>
  </v-col>
</template>

<style lang="scss">
.v-btn:before {
  opacity: 0 !important;
}
.v-ripple__container {
  opacity: 0 !important;
}
</style>

<script>
export default {
  props: ['isExplainClickPoint'],
  data() {
    return {
      buttonLighten: 5,
      buttonLightenString: 'purple lighten-3'
    };
  },
  mounted: async function () {},
  watch: {
    isExplainClickPoint: function (newIsExplainClickPoint, oldIsExplainClickPoint) {
      console.log(newIsExplainClickPoint, oldIsExplainClickPoint, 'dddd');
      if (newIsExplainClickPoint) {
        this.buttonLighten = 5;
        this.buttonLightenString = 'purple lighten-3';
      }
    }
  },
  methods: {
    subButtonLiten() {
      if (this.buttonLighten > 0) {
        this.buttonLighten = this.buttonLighten - 1;
        this.buttonLightenString = `purple lighten-${this.buttonLighten}`;
      } else if (this.buttonLighten === 0) {
        this.buttonLighten = this.buttonLighten - 1;
        this.buttonLightenString = 'amber';
        this.$emit('add-adjustpoint');
      }
    }
  }
};
</script>
