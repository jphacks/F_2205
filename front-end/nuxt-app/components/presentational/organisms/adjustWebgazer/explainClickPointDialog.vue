<template>
  <v-row justify="center">
    <v-dialog v-model="isOpenExplainClickPointDialog" persistent max-width="490">
      <v-card>
        <v-card-title class="text-h5">まずマウスクリックによって大まかな視線の位置を決定します</v-card-title>
        <v-card-text>
          <ul>
            <li>マーカーをみながら5回クリックしてください</li>
            <li>それを時計回りに全てのマーカーに対して行ってください</li>
            <li>最後に中心のマーカーにたいして行ってください</li>
          </ul>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" text @click="closeDialog"> Agree </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  props: ['isOpenExplainClickPointDialog'],
  data() {
    return {
      dialog: true
    };
  },
  mounted: async function () {},
  methods: {
    closeDialog() {
      this.$emit('close-explain-click-pointDialog');
    }
  }
};
</script>
