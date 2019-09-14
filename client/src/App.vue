<template>
  <v-app>
    <VNavBar />
    <v-content>
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
