import axios from 'axios'

const BASE_URI = 'http://localhost:9090'
const client = axios.create({
  baseURL: BASE_URI,
  jason: true
})

const APIClient = {
  async getPost (pID) {
    const postData = await this.perform('get', `/blog/post/${pID}`)
    return postData
  },
  async getComments (pID) {
    const commentsData = await this.perform('get', `/blog/comment/${pID}`)
    return commentsData
  },
  async getAllPosts () {
    const postsData = await this.perform('get', `/blog/post/verified`)
    return postsData
  },
  async getAllUnverifiedPosts () {
    const postsData = await this.perform('get', `/blog/post/unverified`)
    return postsData
  },
  async getUser (uID) {
    const userData = await this.perform('get', `/blog/user/${uID}`)
    return userData
  },
  async getUserPosts (uID) {
    const postsData = await this.perform('get', `/blog/user/${uID}/posts`)
    return postsData
  },
  async getPostTags (pID) {
    const postTags = await this.perform('get', `/blog/post/${pID}/tags`)
    return postTags
  },
  async postNewPost (post) {
    this.perform('post', `/blog/post`, post)
  },
  async updatePost (pID, tags) {
    this.perform('put', `/blog/post/${pID}`, tags)
  },
  async postNewComment (comment, pID) {
    this.perform('post', `/blog/comment/${pID}`, comment)
  },
  async deleteComment (cID) {
    this.perform('delete', `/blog/comment/${cID}`)
  },
  async deletePost (pID) {
    this.perform('delete', `/blog/post/${pID}`)
  },
  async signUp (user) {
    const response = await this.perform('post', `/login/signup`, JSON.stringify(user))
    return response
  },
  async logIn (email, password) {
    const response = await this.perform('get', `/login/${email}/${password}`)
    return response
  },

  async perform (method, resource, data) {
    // Get access token here
    return client({
      method,
      url: resource,
      data
    }).then(req => {
      return req.data
    })
  }
}

export default APIClient
