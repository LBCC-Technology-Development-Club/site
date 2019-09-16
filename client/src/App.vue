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

export default {
  name: 'App',
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
