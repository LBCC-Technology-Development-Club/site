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
                class="hidden-sm-and-down"
              >{{ member.about }}</v-list-item-subtitle>
              <v-list-item-subtitle
                class="hidden-md-and-up"
              >{{ member.headline }}</v-list-item-subtitle>
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
        { id: 0, name: 'Srikar Valluri', picture: '', headline: 'President', about: 'Srikar is our club President. He is active in all teams, and has significant design influence throughout.', teams: ['software', 'admin', 'mechanical'] },
        { id: 1, name: 'Nathan Shaw', picture: '/img/nathan.jpg', headline: 'Founder', about: 'Nathan is our founder. Leading the mechanical design he does the majority of our CAD work.', teams: ['admin', 'mechanical'] },
        { id: 2, name: 'Jonathan Landers', picture: '', headline: 'Active Member', about: 'Jonathan\'s bio.', teams: ['software', 'admin'] },
        { id: 3, name: 'John Warila', picture: 'img/john.jpg', headline: 'Active Member', about: 'John is the primary web-developer - he made this site! - and helps with some of the ECE work for the prosthetic.', teams: ['software', 'admin'] },
        { id: 4, name: 'Steven Nelson', picture: '', headline: 'Active Member', about: 'Steven\'s bio.', teams: ['software', 'admin'] },
        { id: 5, name: 'Stephen Oium', picture: '', headline: 'Active Member', about: 'Stephen\'s bio.', teams: ['software', 'admin'] },
        { id: 6, name: 'Mico Santiago', picture: 'img/mico.jpg', headline: 'Active Member', about: 'Mico is a pre-med student at OSU / LBCC studying Biological Engineering. He focuses on administration and social outreach.', teams: ['admin'] }
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
