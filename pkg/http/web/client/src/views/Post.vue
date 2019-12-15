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
        <router-link to="/blog" id="go-back-link">Back to the feed</router-link>
        <VPostTitle
          v-bind:title="currentPost.title"
          v-bind:author="currentPost.author"
          v-bind:uID="currentPost.user_id"
          v-bind:timestamp="currentPost.timestamp"
        />
        <VPostContent
          v-bind:body="currentPost.body"
        />
        <v-btn
          text
          v-if="isAdmin"
          color="primary"
          :to="{ name: 'editPost', params: { id: currentPost.post_id }}"
        >edit</v-btn>
        <VPostComments/>
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
import VPostTitle from '@/components/VPostTitle.vue'
import VPostContent from '@/components/VPostContent.vue'
import VPostComments from '@/components/VPostComments.vue'
import APIClient from '@/apiClient'

export default {
  components: {
    VPostTitle,
    VPostContent,
    VPostComments
  },
  data () {
    return {
      currentPost: {}
    }
  },
  computed: {
    isAdmin () {
      if (this.getCookie('admin') === 'true') return true
      if (this.getCookie('admin') === 'false') return false
    }
  },
  mounted () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      APIClient.getPost(this.$route.params.id).then(responseJSON => {
        this.currentPost = responseJSON
      })
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
  },
}
</script>

<style>
#go-back-link {
  text-decoration: none;
}
</style>
