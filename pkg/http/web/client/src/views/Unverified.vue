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
  },
  mounted () {
    if (this.getCookie('jwt') === '') {
      this.$router.push({ name: 'blog' })
    }
    if (this.getCookie('admin') === 'false') {
      this.$router.push({ name: 'blog' })
    }
  },
  methods: {
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
