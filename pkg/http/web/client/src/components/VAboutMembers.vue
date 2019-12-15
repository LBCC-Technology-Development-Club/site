<template>
  <v-container>
    <v-list class="grey lighten-5">
      <v-list-item
        class="member-card"
        v-for="(member, index) in members"
        :key="member.id"
      >
        <v-card
          outlined
          elevation="2"
        >
          <v-list-item three-line>
            <v-list-item-avatar
              v-if="index % 2 == 0"
              size="100"
            ><v-img :src="member.picture"></v-img></v-list-item-avatar>
            <v-list-item-content>
              <div class="overline mb-4"><em>teams</em><span v-for="team in member.teams" :key="team"> - {{ team }} </span></div>
              <v-list-item-title class="headline mb-1">{{ member.name }}</v-list-item-title>
              <v-list-item-subtitle
              >{{ member.headline }}</v-list-item-subtitle>
              <v-list-item-content
                class="hidden-sm-and-down"
              >{{ member.about }}</v-list-item-content>
            </v-list-item-content>
            <v-list-item-avatar
              v-if="index % 2 !== 0"
              size="100"
              color="gray"
            ><v-img :src="member.picture"></v-img></v-list-item-avatar>
          </v-list-item>
        </v-card>
      </v-list-item>
    </v-list>
  </v-container>
</template>

<script>
export default {
  name: 'AboutMembers',
  data () {
    return {
      members: [
        { id: 0, name: 'Srikar Valluri', picture: '/img/srikar_opt.png', headline: 'President', about: 'Hiya! My name is Srikar Valluri, and I\'m an Electrical Engineering and Computer Science student at Linn-Benton Community College. I\'m the President of the Technology Development Club, and I specialize in writing the Microcontroller Code and designing the Electrical Circuits for our club. I also am very active in contributing to the design, as it has several various mechanical and electrical components, and I also do a lot of Administration work as president of the club.', teams: ['software', 'admin', 'mechanical'] },
        { id: 1, name: 'Nathan Shaw', picture: '/img/nathan_opt.jpg', headline: 'Founder', about: 'My name is Nathaniel. I founded the club with the goal of bringing something into the world that can help people. I love projects and hard problems. Lead Designer and Club founder.', teams: ['admin', 'mechanical'] },
        { id: 2, name: 'Jonathan Landers', picture: '/img/jonathan_opt.jpg', headline: 'Active Member', about: 'My name is Jonathan Landers, I am a Mechanical Engineering student at OSU/LBCC. I work on everything that needs to be done in mechanical, CAD, design, and 3D printing.', teams: ['mechanical'] },
        { id: 3, name: 'John Warila', picture: 'img/john_opt.jpg', headline: 'Active Member', about: 'I\'m John Warila, a Computer Science Student at OSU/LBCC. I am our primary web-developer and help with some of the ECE and design work for the prosthetic.', teams: ['software', 'admin'] },
        { id: 4, name: 'Mico Santiago', picture: 'img/mico_opt.jpg', headline: 'Active Member', about: 'My name is Mico Santiago, pre-med student at OSU/LBCC where I am studying Biological Engineering. My responsibilities is administration and social outreach. I was a former pro tennis player and I aspire to save lives as a doctor.', teams: ['admin'] }
      ]
    }
  }
}
</script>

<style lang="css">
  .member-card {
    padding-top: 1em;
  }
</style>
