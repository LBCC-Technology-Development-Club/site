<template>
<div id="postcard">
  <v-card>
    <v-card-title>{{ title }}</v-card-title>
    <v-card-text>
      <div>by {{ author }}</div>
      <br>
      <div class="text--primary">{{ summary }}</div>
    </v-card-text>
  </v-card>
</div>
</template>

<script>

export default {
  name: 'VPostCard',
  props: ['title', 'author', 'summary']
}
</script>

<style>
#postcard {
  padding-bottom: 20px;
}
</style>
