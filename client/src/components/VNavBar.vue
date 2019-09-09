<template>
  <v-app-bar
    app
    class="primary"
    collapse-on-scroll
  >
      <v-toolbar-title class="hidden-sm-and-down headline text-uppercase">
        <router-link to="/" id="logo">
          <span class="white-text">LBCC</span>
          <span class="white-text font-weight-light"> - Technology Development Club</span>
        </router-link>
      </v-toolbar-title>
      <v-toolbar-title class="hidden-md-and-up headline text-uppercase">
        <router-link to="/" id="logo">
          <span class="white-text font-weight-light">TDC</span>
        </router-link>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-menu
        bottom
        left
      >
        <template v-slot:activator="{ on }">
          <v-btn
            icon
            color="white"
            v-on="on"
          >
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </template>
        <v-list>
          <v-list-item
            v-for="item in menu"
            :key="item.link"
            v-bind:to="item.link"
          >
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>
</template>

<script>
export default {
  name: 'NavBar',
  data () {
    return {
      menu: [
        { title: 'Home', link: '/' },
        { title: 'About', link: '/about' },
        { title: 'Blog', link: '/blog' }
      ]
    }
  }

}
</script>

<style lang="css">
.white-text {
  color: white;
}
#logo {
  text-decoration: none;
}
</style>
