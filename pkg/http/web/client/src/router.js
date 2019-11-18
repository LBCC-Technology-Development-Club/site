import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import About from './views/About.vue'
import Blog from './views/Blog.vue'
import Post from './views/Post.vue'
import CreatePost from './views/CreatePost.vue'
Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/about',
      name: 'about',
      component: About
    },
    {
      path: '/blog',
      name: 'blog',
      component: Blog
    },
    {
      path: '/blog/post/:id',
      name: 'post',
      component: Post
    },
    {
      path: '/blog/create-post',
      name: 'create-post',
      component: CreatePost
    }
  ]
})
