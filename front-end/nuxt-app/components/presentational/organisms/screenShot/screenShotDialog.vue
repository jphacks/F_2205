<template>
  <div class="text-center">
    <v-dialog v-model="isOpenScreenShotDialog" persistent width="450" content-class="dialogContainer">
      <div class="countContainer">
        <div>{{ this.currentScreenShotCount }}</div>
      </div>
    </v-dialog>
  </div>
</template>

<style>
.countContainer {
  display: flex;
  justify-content: center;
  font-size: 300px;
  opacity: 0.46;
  background-color: #21212182;
  border-color: #21212182;
}
.dialogContainer {
  border-radius: 450px;
}
</style>

<script>
export default {
  props: ['currentScreenShotCount', 'isOpenScreenShotDialog'],
  data() {
    return {
      dialog: false
    };
  },

  watch: {
    isOpenScreenShotDialog: function (newState, oldState) {
      console.log(newState);
    }
  }
};
</script>
