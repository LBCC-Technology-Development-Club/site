<template>
  <v-container>
    <h1>Our Current Project</h1>
    <v-card
      elevation="4"
    >
      <v-img
        class="white--text"
        alt="Image of a mechanical ankle."
        src="@/assets/img/ankleRender.png"
      >
      </v-img>
      <v-card-text>A design prototype of the prosthetic ankle</v-card-text>
    </v-card>
    <br>
    <p>Our current project is to design a low cost, but highly functional, prosthetic for below the knee amputees. By combining an innovative biomimetic design with advanced electronics, we hope to provide a massive increase in mobility for those in need.</p>
    <br>
    <h2>The Brains</h2>
    <v-img
      class="white--text"
      alt="Image of a microcontroller."
      src="@/assets/img/teensy35.jpg"
    >
    </v-img>
    <br>
    <p>This is one of our microcontrollers, these are the brains behind the prosthetic. By melding our mechanical design with cutting edge software and electronics, we are able to squeeze the most performance out of the prosthetic.</p>
    <hr>
    <br>
    <p>For updates on our progress, visit our blog!</p>
    <v-btn
      text
      color="primary"
      to="/blog"
    >Updates</v-btn>
  </v-container>
</template>

<script>
export default {
  name: 'CurrentProject'
}
</script>
