<template>
  <v-app light>
  <div>
    <!-- ヘッダー -->
    <Header />
    <!-- ヘッダー -->
  </div>

    <nuxt />

  </v-app>
</template>

<script>
import Header from '~/components/presentational/organisms/header';

export default {
  components: {
    Header
  },
}
</script>
