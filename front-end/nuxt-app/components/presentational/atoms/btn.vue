<template>
    <v-btn :color="color" @click.native="clickedfn">{{ text }}</v-btn>
</template>

<script>
export default {
    props: ['text', 'color', 'clickedfn']
}
</script>
