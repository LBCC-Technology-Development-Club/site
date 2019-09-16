import Vue from 'vue'
import Vuetify from 'vuetify/lib'

Vue.use(Vuetify)

export default new Vuetify({
  icons: {
    iconfont: 'mdi'
  },
  theme: {
    light: {
      primary: '#3F51B5',
      secondary: '#009688',
      warning: '#F44336'
    }
  }
})
