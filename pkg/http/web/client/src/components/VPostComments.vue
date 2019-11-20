<template>
<div id="post-comments">
  <h2>Comments</h2>
  <VComment
    v-for="comment in comments"
    v-bind:key="comment.commentID"
    v-bind:commenter="comment.commenter"
    v-bind:timestamp="comment.timestamp"
    v-bind:body="comment.body"
  >
  </VComment>
</div>
</template>

<script>
import VComment from '@/components/VComment.vue'

export default {
  name: 'VPostComments',
  components: {
    VComment
  },
  props: ['comments']
}
</script>

<style>
#postcard {
  padding-bottom: 20px;
}
</style>
