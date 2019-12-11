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
        <v-btn
          text
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
  mounted () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      APIClient.getPost(this.$route.params.id).then(responseJSON => {
        this.currentPost = responseJSON
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
