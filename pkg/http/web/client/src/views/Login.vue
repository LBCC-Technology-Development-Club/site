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
        <router-link to="/blog/signup" id="go-back-link">Need an account? Sign Up!</router-link>
        <v-form>
          <v-text-field
            v-model="email"
            :rules="emailRules"
            label="E-mail"
            required
          ></v-text-field>
          <v-text-field
            v-model="password"
              :append-icon="show1 ? 'mdi-eye' : 'mdi-eye-off'"
              :type="show1 ? 'text' : 'password'"
              :rules="passwordRules"
              name="input-10-1"
              label="Password"
              @click:append="show1 = !show1"
          ></v-text-field>
          <v-btn
            :disabled="!valid"
            color="success"
            class="mr-4"
            @click.prevent="getFormValues()"
          >
            Login
          </v-btn>
        </v-form>
        <p v-if="invalid">Invalid Login</p>
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
import { mapGetters, mapActions } from 'vuex'
import { store } from '@/store'
import APIClient from '@/apiClient'

export default {
  components: {
  },
  data () {
    return {
      user: {},
      valid: true,
      email: '',
      emailRules: [
        v => !!v || 'E-mail is required',
        v => /.+@.+\..+/.test(v) || 'E-mail must be valid'
      ],
      passwordRules: [
        v => !!v || 'Password is required'
      ],
      password: '',
      show1: false,
      invalid: false
    }
  },
  computed: {
    ...mapGetters([
      'jwt',
      'isAdmin',
      'userID'
    ])
  },
  methods: {
    getFormValues () {
      this.user.email = this.email
      this.user.password = this.password
      APIClient.logIn(this.user).then(responseJSON => {
        if (responseJSON["message"] === "Invalid login") {
          this.invalid = true
        }
        if (responseJSON["message"] === "Logged in") {
          const jwt = responseJSON["jwt"]
          const userid = responseJSON["user"]
          const admin = responseJSON["admin"]
          document.cookie = "jwt=" + jwt + "; path=/"
          document.cookie = "userid=" + String(userid) + "; path=/"
          document.cookie = "admin=" + admin + "; path=/"
          console.log(this.getCookie('admin'))
          console.log(this.getCookie('jwt'))
          console.log(this.getCookie('userid'))
          this.$router.push({ name: 'blog' })
        }
      })
    },
    ...mapActions([
      `setUser`
    ]),
    getCookie (name) {
      const nameEQ = name + "="
      var ca = document.cookie.split(';')
      for (var i = 0; i < ca.length; i++) {
        var c = ca[i]
        while (c.charAt(0)===' ') c = c.substring(1, c.length)
        if (c.indexOf(nameEQ) === 0) return c.substring(nameEQ.length, c.length)
      }
      console.log('Bad cookie name')
      return null
    }
  }
}
</script>

<style>
#go-back-link {
  text-decoration: none;
}
</style>
