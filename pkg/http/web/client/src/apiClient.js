import axios from 'axios'

const BASE_URI = 'http://lbcctechdevclub.org'
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
    const postsData = await this.perform('get', `/blog/post`)
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
