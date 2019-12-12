<template>
<div id="post-comments">
  <h2>Comments</h2>
  <div v-if="isLoggedIn">
    <v-row>
      <v-col cols="12" sm="8">
        <v-text-field label="Comment" single-line clearable v-model="content"></v-text-field>
      </v-col>
      <v-col cols="12" sm="4">
        <v-btn color="primary" text @click.prevent="getFormValues()">submit</v-btn>
      </v-col>
    </v-row>
  </div>
  <VComment
    v-for="comment in comments"
    v-bind:key="comment.comment_id"
    v-bind:cID="comment.comment_id"
    v-bind:commenter="comment.author"
    v-bind:timestamp="comment.timestamp"
    v-bind:body="comment.content"
  >
  </VComment>
</div>
</template>

<script>
import VComment from '@/components/VComment.vue'
import APIClient from '@/apiClient'

export default {
  name: 'VPostComments',
  components: {
    VComment
  },
  data () {
    return {
      comments: {},
      comment: {
        user_id: 0,
        pID: 0,
        content: '',
        timestamp: ''
      },
      content: ''
    }
  },
  mounted () {
    this.fetchData()
  },
  computed: {
    isLoggedIn () {
      if (this.getCookie('jwt') === '') return false
      return true
    }
  },
  methods: {
    fetchData () {
      APIClient.getComments(this.$route.params.id).then(responseJSON => {
        this.comments = responseJSON
      })
    },
    getFormValues () {
      this.comment.content = this.content
      this.comment.user_id = parseInt(this.getCookie('userid'))

      APIClient.postNewComment(this.comment, this.$route.params.id)
      this.comments = {}
      setTimeout(() => this.fetchData(), 100)
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

<style>
#postcard {
  padding-bottom: 20px;
}
</style>
