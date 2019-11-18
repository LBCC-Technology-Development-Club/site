<template>
  <p>{{ currentPost.title }}</p>
</template>

<script>

export default {
  components: {
  },
  data () {
    return {
      posts: [
        { pid: 0, title: 'Our New Blog!', author: 'John Warila', timestamp: 'Nov 17, 2019', summary: 'This article talks about our new blog, and the tech inside.' },
        { pid: 1, title: 'Weekly Update: Nov 17 2019', author: 'Srikar Valluri', timestamp: 'Nov 17, 2019', summary: 'The weekly update from the week of November 17, 2019. Includes team progress and our meeting minutes.' }
      ],
      currentPost: {}
    }
  },
  created () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      this.currentPost = this.posts[this.$route.params.id]
    }
  }
}
</script>
