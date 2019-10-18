<template>
  <v-container>
    <h1>Our Current Project</h1>
    <br>
    <h3>Low Cost Biomimetic Prosthesis</h3>
    <v-carousel
      cycle
      hide-delimiter-background
      show-arrows-on-hover
    >
      <v-carousel-item
        v-for="(image, i) in images"
        :key="i"
        :src="image.src"
      >
      </v-carousel-item>
    </v-carousel>
    <br>
    <p>Our current project is to design a low cost, but highly functional, prosthetic for below the knee amputees. By combining an innovative biomimetic design with advanced electronics, we hope to provide a massive increase in mobility for those in need.</p>
    <br>
    <h3>For updates on our progress, visit our blog!</h3>
    <v-row
      align="center"
      justify="center"
    >
      <v-btn
        text
        large
        color="primary"
        to="/blog"
      >Updates</v-btn>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'CurrentProject',
  data () {
    return {
      images: [
        { src: 'img/prototype-design-0_opt.png' },
        { src: 'img/prototype-design-1_opt.png' },
        { src: 'img/prototype-design-2_opt.png' },
        { src: 'img/prototype-design-3_opt.png' }
      ]
    }
  }
}
</script>
