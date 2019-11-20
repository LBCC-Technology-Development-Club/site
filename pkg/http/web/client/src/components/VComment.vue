<template>
<div id="comment">
  <v-card
    outlined
  >
    <v-card-text>{{ commenter }}, {{ timestamp }} <br> <span class="text--primary">{{ body }}</span></v-card-text>
  </v-card>
</div>
</template>

<script>

export default {
  name: 'VComment',
  props: ['commenter', 'timestamp', 'body']
}
</script>

<style>
#comment {
  padding-bottom: 20px;
}
</style>
