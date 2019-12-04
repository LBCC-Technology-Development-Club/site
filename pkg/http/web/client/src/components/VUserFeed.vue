<template>
<div id="feed">
    <VPostCard
      v-for="post in posts"
      v-bind:key="post.post_id"
      v-bind:pid="post.post_id"
      v-bind:title="post.title"
      v-bind:author="post.author"
      v-bind:timestamp="post.timestamp"
      v-bind:summary="post.summary"
    >
  </VPostCard>
</div>
</template>

<script>
import VPostCard from '@/components/VPostCard.vue'
import APIClient from '@/apiClient'

export default {
  name: 'Feed',
  components: { VPostCard },
  data () {
    return {
      posts: []
    }
  },
  created () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      APIClient.getUserPosts(this.$route.params.id).then(responseJSON => {
        this.posts = responseJSON
      })
    }
  }
}
</script>
