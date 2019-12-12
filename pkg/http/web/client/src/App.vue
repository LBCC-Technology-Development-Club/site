<template>
  <v-app>
    <VNavBar />
    <v-content class="no-pad">
      <router-view></router-view>
    </v-content>
  </v-app>
</template>

<script>
import VNavBar from '@/components/VNavBar.vue'

import { store } from '@/store'

export default {
  name: 'App',
  store,
  components: {
    VNavBar
  },
  data: () => ({
    //
  })
}
</script>

<style lang="css">
  .no-pad {
    padding: 0;
  }
</style>
