<template>
  <v-form>
    <v-text-field label="Title" single-line v-model="title"></v-text-field>
    <v-text-field label="Summary" single-line v-model="summary"></v-text-field>
    <v-textarea label="Body" auto-grow v-model="body"></v-textarea>
    <v-btn
      color="primary"
      @click.prevent="getFormValues()"
      >POST</v-btn>
  </v-form>
</template>

<script>
import APIClient from '@/apiClient'

export default {
  name: 'VPostInput',
  data () {
    return {
      post: {
        title: '',
        summary: '',
        author: '',
        body: '',
        timestamp: ''
      },
      title: '',
      summary: '',
      body: ''
    }
  },
  methods: {
    getFormValues () {
      this.post.title = this.title
      this.post.summary = this.summary
      this.post.body = this.body
      this.author = 'Admin Test User'

      APIClient.postNewPost(this.post)

      this.$router.push({ name: 'blog' })
    }
  }
}
</script>
