<template>
<div id="postcard">
  <v-card>
    <v-card-title>{{ title }}</v-card-title>
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

export default {
  name: 'VPostCard',
  props: ['pid', 'title', 'author', 'timestamp', 'summary']
}
</script>

<style>
#postcard {
  padding-bottom: 20px;
}
</style>
