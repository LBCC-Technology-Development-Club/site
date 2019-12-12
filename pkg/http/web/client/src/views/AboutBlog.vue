<template name="component-name">
  <v-container>
    <v-row>
      <v-col
        cols="12"
        md="2"
      >
      </v-col>
      <v-col
        cols="12"
        md="8"
      >
        <h1>About Page</h1>
        <p>The blog focuses around the 'feed'. This page shows all the posts that have been verified, with any stickied posts on top</p>
        <router-link to="/blog" id="go-back-link">Go to the feed</router-link>
        <p>Lists all posts</p>
        <router-link to="/blog/login" id="go-back-link">Go to login</router-link>
        <p>Lets you log in</p>
        <router-link to="/blog/unverified" id="go-back-link">Go to the unverified feed</router-link>
        <p>Shows the admin only unverified feed</p>
      </v-col>
      <v-col
        cols="12"
        md="2"
      >
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  components: {
  }
}
</script>
