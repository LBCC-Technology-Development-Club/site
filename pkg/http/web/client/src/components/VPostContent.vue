<template>
<div id="post-content">
  <p>{{ body }}</p>
</div>
</template>

<script>

export default {
  name: 'VPostContent',
  props: ['body']
}
</script>

<style>
#postcard {
  padding-bottom: 20px;
}
</style>
