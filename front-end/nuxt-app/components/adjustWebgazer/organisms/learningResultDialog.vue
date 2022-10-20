<template>
  <v-row justify="center">
    <v-dialog v-model="this.isOpenLearningResultDialog" persistent max-width="390">
      <v-card>
        <v-card-title class="text-h5">結果が出ました</v-card-title>
        <v-card-text>現在の精度: {{ this.recisionMeasurement }}%</v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" text @click="closeDialog"> Disagree </v-btn>
          <v-btn color="green darken-1" text @click="closeDialog"> Agree </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  props: ['isOpenLearningResultDialog', 'recisionMeasurement'],
  data() {
    return {
      dialog: false
    };
  },
  methods: {
    closeDialog() {
      this.$emit('close-learningresultdialog');
    }
  }
};
</script>
