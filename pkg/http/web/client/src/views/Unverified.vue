<template>
  <v-container>
    <v-row>
      <v-col
        cols="12"
        md="2"
      >
      </v-col>
      <v-col
        cols="12"
        md="8"
      >
        <h1>TDC Blog</h1>
        <VUnverifiedFeed />
      </v-col>
      <v-col
        cols="12"
        md="2"
      >
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import VUnverifiedFeed from '@/components/VUnverifiedFeed.vue'

export default {
  components: {
    VUnverifiedFeed
  }
}
</script>
