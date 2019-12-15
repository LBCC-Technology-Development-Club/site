<template>
<div id="postcard">
  <v-card>
    <span v-if="inArray('sticky', tags)"><v-icon>mdi-pin</v-icon></span>
    <span v-if="inArray('update', tags)"><v-icon color="primary">mdi-update</v-icon></span>
    <span v-if="inArray('announcement', tags)"><v-icon color="red">mdi-bullhorn</v-icon></span>
    <br>
    <v-card-title>
      {{ title }}
    </v-card-title>
    <v-card-text>
      <div>by {{ author }}, {{ timestamp }}</div>
      <br>
      <div class="text--primary">{{ summary }}</div>
    </v-card-text>
    <v-btn
      text
      color="primary"
      :to="{ name: 'post', params: { id: pid }}"
    >Read</v-btn>
  </v-card>
</div>
</template>

<script>
import APIClient from '@/apiClient'

export default {
  name: 'VPostCard',
  props: ['pid', 'title', 'author', 'timestamp', 'summary'],
  data () {
    return {
      tags: []
    }
  },
  mounted () {
    this.fetchData()
  },
  methods: {
    inArray (value, array) {
      const length = array.length
      for (var i = 0; i < length; i++) {
        if (array[i] === value) return true
      }
      return false
    },
    fetchData () {
      APIClient.getPostTags(this.pid).then(responseJSON => {
        this.tags = responseJSON
      })
    }
  }
}
</script>

<style>
#postcard {
  padding-bottom: 20px;
}
</style>
