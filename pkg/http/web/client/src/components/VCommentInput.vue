<template>
<div id="post-title">
  <v-row>
    <v-col cols="12" sm="8">
      <v-text-field label="Comment" single-line></v-text-field>
    </v-col>
    <v-col cols="12" sm="4">
      <v-btn color="primary" text>submit</v-btn>
    </v-col>
  </v-row>
</div>
</template>

<script>

export default {
  name: 'VCommentInput'
}
</script>

<style>
#postcard {
  padding-bottom: 20px;
}
</style>
