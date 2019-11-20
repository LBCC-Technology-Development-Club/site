<template>
  <v-container>
    <v-row>
      <v-col
        cols="12"
        md="3"
      >
      </v-col>
      <v-col
        cols="12"
        md="6"
      >
        <router-link to="/blog" id="go-back-link">Back to the feed</router-link>
        <VPostTitle
          v-bind:title="currentPost.title"
          v-bind:author="currentPost.author"
        />
        <VPostContent
          v-bind:body="currentPost.body"
        />
        <VPostComments
          v-bind:comments="currentPost.comments"
        />
      </v-col>
      <v-col
        cols="12"
        md="3"
      >
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import VPostTitle from '@/components/VPostTitle.vue'
import VPostContent from '@/components/VPostContent.vue'
import VPostComments from '@/components/VPostComments.vue'

export default {
  components: {
    VPostTitle,
    VPostContent,
    VPostComments
  },
  data () {
    return {
      posts: [
        {
          pid: 0,
          title: 'Our New Blog!',
          author: 'John Warila',
          timestamp: 'Nov 17, 2019',
          summary: 'This article talks about our new blog, and the tech inside.',
          body: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas venenatis porttitor dolor, pellentesque convallis ipsum luctus non. Nullam posuere tellus id quam placerat dapibus. Duis at porta dolor, ultrices vulputate turpis. Ut nec metus sed dui laoreet sagittis. Quisque augue est, ullamcorper ut metus nec, lobortis tincidunt turpis. Phasellus ornare elit vitae felis vulputate, eu ornare ex faucibus. Nam elementum turpis in convallis tincidunt. Duis vel ligula quis velit posuere tincidunt tempor in dui.Sed scelerisque erat sit amet purus lobortis mollis. Praesent ac mi eget mauris interdum eleifend. Curabitur posuere lectus vitae gravida iaculis. Vivamus id bibendum mi. Maecenas in quam risus. Cras eleifend nibh est, ac lacinia nibh euismod sed. Donec vel ornare lacus. Maecenas scelerisque elementum magna, id porta felis condimentum eget. Nunc sed aliquet lorem. Morbi sagittis vestibulum lorem in facilisis. Ut lorem ex, egestas a tellus sit amet, ornare rutrum arcu.',
          comments: [
            {
              commenter: 'John Warila',
              body: 'This is a comment!',
              timestamp: '5:56pm Nov 19',
              commentID: 0
            },
            {
              commenter: 'Srikar Valluri',
              body: 'Cool!',
              timestamp: '5:56pm Nov 19',
              commentID: 1
            }
          ]
        },
        {
          pid: 1,
          title: 'Weekly Update: Nov 17 2019',
          author: 'Srikar Valluri',
          timestamp: 'Nov 17, 2019',
          summary: 'The weekly update from the week of November 17, 2019. Includes team progress and our meeting minutes.',
          body: 'Update',
          comments: [
            {
              commenter: 'John Warila',
              body: 'This is a comment!',
              timestamp: '5:56pm Nov 19',
              commentID: 2
            },
            {
              commenter: 'Srikar Valluri',
              body: 'Cool!',
              timestamp: '5:56pm Nov 19',
              commentID: 3
            }
          ]
        }
      ],
      currentPost: {}
    }
  },
  created () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      this.currentPost = this.posts[this.$route.params.id]
    }
  }
}
</script>

<style>
#go-back-link {
  text-decoration: none;
}
</style>