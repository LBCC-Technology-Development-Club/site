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
        <v-form>
          <v-text-field
            v-model="email"
            :rules="emailRules"
            label="E-mail"
            required
          ></v-text-field>
          <v-text-field
            v-model="name"
            :rules="nameRules"
            label="Name"
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
            color="success"
            class="mr-4"
            @click.prevent="getFormValues()"
          >
            SignUp
          </v-btn>
        </v-form>
        <p v-if="alreadyExist">A user with this email already exists</p>
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
import APIClient from '@/apiClient'

export default {
  components: {
  },
  data () {
    return {
      user: {
        
      },
      name: '',
      nameRules: [
        v => !!v || 'Name is required',
        v => (v && v.length <= 20) || 'Name must be less than 20 characters'
      ],
      email: '',
      emailRules: [
        v => !!v || 'E-mail is required',
        v => /.+@.+\..+/.test(v) || 'E-mail must be valid'
      ],
      password: '',
      passwordRules: [
        v => !!v || 'Password is required'
      ],
      show1: false,
      alreadyExist: false
    }
  },
  methods: {
    getFormValues () {
      const user = {
        uID: '',
        name: this.name,
        email: this.email,
        isAdmin: false,
        hash: '',
        password: this.password,
        role: 'member'
      }
      APIClient.signUp(user).then(responseJSON => {
        if (responseJSON["message"] === "A user with this email already exists") {
          this.alreadyExist = true
        }
        if (responseJSON["message"] === "Successfuly signed up") {
          // We should get a jwt now
          
        }
      })
    }
  }
}
</script>
