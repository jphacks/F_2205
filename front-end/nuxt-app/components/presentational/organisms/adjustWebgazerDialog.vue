<template>
  <v-row justify="center">
    <v-dialog v-model="isOpenAdjustWebGazerDialog" fullscreen>
      <v-card>
        <v-toolbar>
          <v-btn icon dark @click="handleAdjustWebGazer">
            <v-icon>mdi-close</v-icon>
          </v-btn>
          <v-toolbar-title>Settings</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-toolbar-items>
            <v-btn dark text @click="handleAdjustWebGazer"> Save </v-btn>
          </v-toolbar-items>
        </v-toolbar>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  props: ['isOpenAdjustWebGazerDialog', 'handleAdjustWebGazer']
};
</script>
