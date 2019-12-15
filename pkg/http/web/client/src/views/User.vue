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
        <h1>{{ user.name }}'s Posts</h1>
        <v-btn v-if="isAdmin && !userIsAlreadyAdmin" color="primary" text @click.prevent="makeAdmin()">Make Admin</v-btn>
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
  computed: {
    isAdmin () {
      if (this.getCookie('admin') === 'true') return true
      return false
    },
    userIsAlreadyAdmin () {
      if (this.user.is_admin === true) return true
      return false
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
    },
    makeAdmin () {
      APIClient.makeUserAdmin(this.$route.params.id)
    },
    getCookie (name) {
      const nameEQ = name + "="
      var ca = document.cookie.split(';')
      for (var i = 0; i < ca.length; i++) {
        var c = ca[i]
        while (c.charAt(0)===' ') c = c.substring(1, c.length)
        if (c.indexOf(nameEQ) === 0) return c.substring(nameEQ.length, c.length)
      }
      return null
    }
  }
}
</script>
