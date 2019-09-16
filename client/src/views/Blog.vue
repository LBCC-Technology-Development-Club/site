<template>
  <v-container>
    <v-row>
      <v-col
        cols="12"
        md="2"
      >
      </v-col>
      <v-col
        cols="12"
        md="8"
      >
        <VBlogPlaceHolder />
      </v-col>
      <v-col
        cols="12"
        md="2"
      >
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import VBlogPlaceHolder from '@/components/VBlogPlaceHolder.vue'

export default {
  components: {
    VBlogPlaceHolder
  }
}
</script>
