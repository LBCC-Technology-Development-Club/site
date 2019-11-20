<template>
  <v-container>
    <v-row>
      <v-col
        cols="12"
        md="3"
      >
      </v-col>
      <v-col
        cols="12"
        md="6"
      >
        <h1>{{ users[this.$route.params.id] }}'s Posts</h1>
        <VUserFeed
          v-bind:user="users[this.$route.params.id]"
        />
      </v-col>
      <v-col
        cols="12"
        md="3"
      >
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import VUserFeed from '@/components/VUserFeed.vue'

export default {
  components: {
    VUserFeed
  },
  data () {
    return {
      users: ["John Warila", "Srikar Valluri"]
    }
  }
}
</script>