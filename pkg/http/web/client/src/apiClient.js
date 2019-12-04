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
  getComments (pID) {
    return this.perform('get', `/blog/comments/${pID}`)
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
