<template>
<div id="post-title">
  <h1>{{ title }}</h1>
  <span>
    <router-link :to="{ name: 'user', params: { id: uID } }" id="go-back-link">by {{ author }}, </router-link>
    <p>at {{ timestamp }}</p>
  </span>
</div>
</template>

<script>

export default {
  name: 'VPostTitle',
  props: ['title', 'author', 'timestamp', 'uID']
}
</script>

<style>
#postcard {
  padding-bottom: 20px;
}
</style>
