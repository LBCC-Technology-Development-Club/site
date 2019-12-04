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
          v-bind:timestamp="currentPost.timestamp"
        />
        <VPostContent
          v-bind:body="currentPost.body"
        />
        <VPostComments
          v-bind:comments="comments"
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
      comments: {},
      currentPost: {}
    }
  },
  created () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      console.log()
      APIClient.getPost(this.$route.params.id).then((responseJSON) => {
        this.currentPost = responseJSON
      })
      APIClient.getComments(this.$route.params.id).then(responseJSON => {
        this.comments = responseJSON
      })
    }
  }
}
</script>

<style>
#go-back-link {
  text-decoration: none;
}
</style>
