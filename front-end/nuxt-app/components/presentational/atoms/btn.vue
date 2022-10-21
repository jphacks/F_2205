<template>
  <v-btn class="pl-10 pr-10" :color="color" @click.native="clickedfn">{{ text }}</v-btn>
</template>

<script>
export default {
  props: ['text', 'color', 'clickedfn']
};
</script>
