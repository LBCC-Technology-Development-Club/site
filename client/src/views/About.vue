<template>
  <v-container class="no-pad grey lighten-5">
    <v-row>
      <v-col
      cols="12"
      md="1"
      >
      </v-col>
      <v-col
      cols="12"
      md="4"
      >
        <VAboutText />
      </v-col>
      <v-col
      cols="12"
      md="6"
      >
        <VAboutMembers />
      </v-col>

      <v-col
      cols="12"
      md="1"
      >
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import VAboutText from '@/components/VAboutText.vue'
import VAboutMembers from '@/components/VAboutMembers.vue'

export default {
  components: {
    VAboutText,
    VAboutMembers
  }
}
</script>
