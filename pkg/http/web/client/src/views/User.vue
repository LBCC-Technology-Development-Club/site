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
        <h1>{{ user.Name }}'s Posts</h1>
        <VUserFeed
          v-bind:user="user"
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
import APIClient from '@/apiClient'

export default {
  components: {
    VUserFeed
  },
  data () {
    return {
      user: {}
    }
  },
  mounted () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      APIClient.getUser(this.$route.params.id).then(responseJSON => {
        this.user = responseJSON
      })
    }
  }
}
</script>
