<template>
<div id="comment">
  <v-card
    v-if="!deleted"
    outlined
  >
    <v-card-text>{{ commenter }}, at {{ timestamp }} <br> <span class="text--primary">{{ body }}</span></v-card-text>
    <v-btn small text color="red" @click.prevent="deleteComment()">delete</v-btn>
  </v-card>
</div>
</template>

<script>
import APIClient from '@/apiClient'

export default {
  name: 'VComment',
  props: ['cID', 'commenter', 'timestamp', 'body'],
  data () {
    return {
      deleted: false
    }
  },
  methods: {
    deleteComment () {
      this.deleted = true
      APIClient.deleteComment(this.cID)
    }
  }
}
</script>

<style>
#comment {
  padding-bottom: 20px;
}
</style>
