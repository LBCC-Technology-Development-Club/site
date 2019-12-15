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
        <span>
          <v-btn v-if="isLoggedIn" color="primary" text @click.prevent="signOut()">sign out</v-btn>
          <v-btn v-if="isAdmin" color="primary" text :to="{ name: 'unverified'}">unverified posts</v-btn>
        </span>
        <VFeed />
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
import VFeed from '@/components/VFeed.vue'

export default {
  components: {
    VFeed
  },
  computed: {
    isLoggedIn () {
      if (this.getCookie('jwt') === '') return false
      return true
    },
    isAdmin() {
      if (this.getCookie('admin') === 'true') return true
      return false
    }
  },
  methods: {
    signOut () {
      document.cookie = "jwt=; path=/"
      document.cookie = "user=; path=/"
      document.cookie = "admin=; path=/"
      this.$router.push({ name: 'signout' })
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
