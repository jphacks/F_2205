<template>
  <v-row justify="center">
    <v-dialog v-model="this.isOpenDialog" persistent max-width="390">
      <v-card>
        <v-card-title class="text-h5">細かな設定をします</v-card-title>
        <v-card-text> センターのマーカーを5秒間見つめてください </v-card-text>
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
  props: ['isOpenDialog'],
  data() {
    return {
      dialog: false
    };
  },
  methods: {
    closeDialog() {
      this.$emit('close-dialog');
    }
  }
};
</script>
