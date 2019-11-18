<template>
<div id="feed">
    <VPostCard
      v-for="post in posts"
      v-bind:key="post.pid"
      v-bind:pid="post.pid"
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

export default {
  name: 'Feed',
  components: { VPostCard },
  data () {
    return {
      posts: [
        { pid: 0, title: 'Our New Blog!', author: 'John Warila', timestamp: 'Nov 17, 2019', summary: 'This article talks about our new blog, and the tech inside.' },
        { pid: 1, title: 'Weekly Update: Nov 17 2019', author: 'Srikar Valluri', timestamp: 'Nov 17, 2019', summary: 'The weekly update from the week of November 17, 2019. Includes team progress and our meeting minutes.' }
      ]
    }
  }
}
</script>
