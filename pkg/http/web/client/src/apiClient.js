import axios from 'axios'

const BASE_URI = 'http://localhost:9090'
const client = axios.create({
  baseURL: BASE_URI,
  jason: true
})

const APIClient = {
  async getPost (pID) {
    let postData = await this.perform('get', `/blog/post/${pID}`)
    return postData
  },
  async getComments (pID) {
    let commentsData = await this.perform('get', `/blog/comment/${pID}`)
    return commentsData
  },
  async getAllPosts () {
    let postsData = await this.perform('get', `/blog/post`)
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
