<template>
  <div>
    <v-text-field label="Title" single-line></v-text-field>
    <v-text-field label="Summary" single-line></v-text-field>
    <v-textarea label="Body" auto-grow></v-textarea>
    <v-btn color="primary">POST</v-btn>
  </div>
</template>

<script>
export default {
  name: 'VPostInput'
}
</script>
