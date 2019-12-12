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
import { mapGetters, mapActions } from 'vuex';

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
      'jwt'
    ])
  },
  methods: {
    getFormValues () {
      APIClient.logIn(this.email, this.password).then(responseJSON => {
        if (responseJSON["message"] === "Invalid login") {
          this.invalid = true
        }
        if (responseJSON["message"] === "Logged in") {
          // do the logged in stuff with a jwt
          this.setJWT(responseJSON["jwt"])
          if (responseJSON["admin"] === "true") this.setAdmin(true)
          if (responseJSON["admin"] === "false") this.setAdmin(false)
          this.$router.push({ name: 'blog' })
        }
      })
    },
    ...mapActions([
      `setJWT`,
      `setAdmin`
    ])
  }
}
</script>

<style>
#go-back-link {
  text-decoration: none;
}
</style>
