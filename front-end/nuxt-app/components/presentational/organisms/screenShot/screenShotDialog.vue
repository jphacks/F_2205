<template>
  <div class="text-center">
    <v-dialog v-model="dialog" width="450" content-class="dialogContainer">
      <template v-slot:activator="{ on, attrs }">
        <v-btn color="red lighten-2" dark v-bind="attrs" v-on="on"> Click Me </v-btn>
      </template>
      <div class="countContainer">
        <div>3</div>
      </div>
    </v-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      dialog: false
    };
  },
  methods: {}
};
</script>
