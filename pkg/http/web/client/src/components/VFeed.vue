<template>
<div id="feed">
    <VPostCard
    v-for="post in posts"
    v-bind:key="post.pid"
    v-bind:title="post.title"
    v-bind:author="post.author"
    v-bind:summary="post.summary">
  </VPostCard>
</div>
</template>

<script>
import VPostCard from '@/components/VPostCard.vue'

export default {
  name: 'Feed',
  components: { VPostCard },
  data () {
    return {
      posts: [
        { pid: 0, title: 'First Post', author: 'John Warila', summary: 'This is the fisrt post' },
        { pid: 1, title: 'Second Post', author: 'Srikar Valluri', summary: 'This is another post' }
      ]
    }
  }
}
</script>
