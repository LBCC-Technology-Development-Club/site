<template>
  <v-container class="grey lighten-5">
    <v-row>
      <v-col
        cols="12"
        md="4"
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
        md="4"
      >
        <v-card
          class="pa-2"
          raised
          title
        >
          <VCurrentProject />
        </v-card>
      </v-col>
      <v-col
        cols="12"
        md="4"
      >
        <v-card
          class="pa-2"
          raised
          title
        >
          <VHomeAbout />
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import VHomeWelcome from '@/components/VHomeWelcome.vue'
import VCurrentProject from '@/components/VCurrentProject.vue'
import VHomeAbout from '@/components/VHomeAbout.vue'

export default {
  components: {
    VHomeWelcome,
    VCurrentProject,
    VHomeAbout
  }
}
</script>
