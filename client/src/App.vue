<template>
  <v-app>
    <v-app-bar app>
      <v-toolbar-title class="headline text-uppercase">
        <span>LBCC</span>
        <router-link to="/" id="logo"><span class="font-weight-light"> - Technology Development Club</span></router-link>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn
        text
        to="/about"
      >
        <span class="mr-2">About</span>
      </v-btn>
      <v-btn
        text
        to="/blog"
      >
        <span class="mr-2">Blog</span>
      </v-btn>
    </v-app-bar>

    <v-content>
      <router-view></router-view>
    </v-content>
  </v-app>
</template>

<script>

export default {
  name: 'App',
  components: {
  },
  data: () => ({
    //
  })
}
</script>

<style lang="css">
#logo {
  text-decoration: none;
  color: black;
}
</style>
