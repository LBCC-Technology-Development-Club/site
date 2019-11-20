<template>
<div id="post-title">
  <h1>{{ title }}</h1>
  <p>by {{ author }}</p>
</div>
</template>

<script>

export default {
  name: 'VPostTitle',
  props: ['title', 'author']
}
</script>

<style>
#postcard {
  padding-bottom: 20px;
}
</style>
