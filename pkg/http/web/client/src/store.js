import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    User: {
      state: {
        currentJWT: '',
        isAdmin: false
      },
      getters: {
        jwt: state => state.currentJWT,
        jwtData: (state, getters) => state.currentJWT ? JSON.parse(atob(getters.jwt.split('.')[1])) : null,
        jwtSubject: (state, getters) => getters.jwtData ? getters.jwtData.sub : null,
        jwtIssuer: (state, getters) => getters.jwtData ? getters.jwtData.iss : null,
        isAdmin: state => state.isAdmin
      },
      mutations: {
        setJWT (state, jwt) {
          state.currentJWT = jwt
        },
        setAdmin (state, admin) {
          state.isAdmin = admin
        }
      },
      actions: {
        fetchJWT ({ commit }, { jwt, admin }) {
          commit('setJWT', jwt)
          commit('setAdmin', admin)
        }
      }
    }
  }
})
