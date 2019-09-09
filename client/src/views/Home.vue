<template>
  <v-container class="grey lighten-5">
    <v-row>
      <v-col
        cols="12"
        sm="4"
      >
        <v-card
          class="pa-2"
          raised
          title
        >
          <VHomeWelcome />
        </v-card>
      </v-col>
      <v-col
        cols="12"
        sm="4"
      >
        <v-card
          class="pa-2"
          raised
          title
        >
          Current Project Component
        </v-card>
      </v-col>
      <v-col
        cols="12"
        sm="4"
      >
        <v-card
          class="pa-2"
          raised
          title
        >
          About Team Component
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import VHomeWelcome from '@/components/VHomeWelcome.vue'
export default {
  components: {
    VHomeWelcome
  }
}
</script>
