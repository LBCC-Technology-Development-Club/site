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
        body: '',
        timestamp: '',
        user_id: 0
      },
      title: '',
      summary: '',
      body: ''
    }
  },
  mounted () {
    if (this.getCookie('jwt') === '') {
      this.$router.push({ name: 'blog' })
    }
  },
  methods: {
    getFormValues () {
      this.post.title = this.title
      this.post.summary = this.summary
      this.post.body = this.body
      
      const userID = parseInt(this.getCookie('userid'))
      console.log(userID)
      this.post.user_id = userID

      const today = new Date()
      const date = today.getFullYear()+'-'+(today.getMonth()+1)+'-'+today.getDate()
      const time = today.getHours()+':'+today.getMinutes()+':'+today.getSeconds()
      const timestamp = date+' '+time

      this.post.timestamp = timestamp
      APIClient.postNewPost(this.post)

      this.$router.push({ name: 'blog' })
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
