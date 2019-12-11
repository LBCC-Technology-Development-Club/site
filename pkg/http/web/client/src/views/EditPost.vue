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
        <h1>Editing:</h1>
        <v-btn
          color="red"
          @click.prevent="deletePost()"
        >delete</v-btn>
        <VPostTitle
          v-bind:title="currentPost.title"
          v-bind:author="currentPost.author"
          v-bind:timestamp="currentPost.timestamp"
        />
        <VPostContent
          v-bind:body="currentPost.body"
        />
        <h2>Tags</h2>
         <v-form>
          <v-checkbox
            v-for="tag in tags"
            :key="tag"
            :label="tag"
            v-model="currentPost.tags"
            :value="tag"
          ></v-checkbox>
          <v-btn
            color="success"
            @click.prevent="getFormValues()"
          >submit</v-btn>
        </v-form>
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
import APIClient from '@/apiClient'

export default {
  components: {
    VPostTitle,
    VPostContent
  },
  data () {
    return {
      currentPost: {},
      tags: [],
      oldTags: []
    }
  },
  mounted () {
    this.fetchData()
  },
  methods: {
    fetchData () {
      APIClient.getPost(this.$route.params.id).then(responseJSON => {
        this.currentPost = responseJSON
        this.oldTags = this.currentPost.tags
      })

      this.tags = ['announcement', 'update', 'verified']
    },
    getFormValues () {
      APIClient.updatePost(this.$route.params.id, { oldTags: this.oldTags, newTags: this.currentPost.tags })
      this.$router.push({ name: 'post', params: { id: this.currentPost.post_id } })
    },
    deletePost () {
      APIClient.deletePost(this.$route.params.id)
      this.$router.push({ name: 'delete' })
    },
    inArray (value, array) {
      const length = array.length
      for (var i = 0; i < length; i++) {
        if (array[i] === value) return true
      }
      return false
    }
  }
}
</script>

<style>
#go-back-link {
  text-decoration: none;
}
</style>
