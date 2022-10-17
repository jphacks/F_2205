<template>
    <v-app-bar
    color="white"
    elevation="10"
    class="header"
    >
        <v-toolbar-title><Logo /></v-toolbar-title>
    </v-app-bar>
</template>

<script>
import Logo from '~/components/presentational/atoms/logo';

export default {
    components: {
        Logo
    },
}
</script>

<style lang="scss">
.header {
    color: #ff9800 !important;
}
</style>
