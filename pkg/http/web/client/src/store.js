import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    currentJWT: '',
    isAdmin: false,
    currentUserID: ''
  },
  getters: {
    jwt: state => state.currentJWT,
    jwtData: (state, getters) => state.currentJWT ? JSON.parse(atob(getters.jwt.split('.')[1])) : null,
    jwtSubject: (state, getters) => getters.jwtData ? getters.jwtData.sub : null,
    jwtIssuer: (state, getters) => getters.jwtData ? getters.jwtData.iss : null,
    isAdmin: state => state.isAdmin,
    userID: state => state.currentUserID
  },
  mutations: {
    setJWT (state, jwt) {
      state.currentJWT = jwt
    },
    setAdmin (state, admin) {
      state.isAdmin = admin
    },
    setUserID (state, userID) {
      state.currentUserID = userID
    }
  },
  actions: {
    setUser ({ commit }, { jwt, userID, admin }) {
      console.log(jwt)
      commit('setJWT', jwt)
      console.log(admin)
      commit('setAdmin', admin)
      console.log(userID)
      commit('setUserID', userID)
    }
  }
})
