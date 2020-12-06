<template>
  <section class="hero is-fullheight">
   <div class="hero-body">
      <div class="columns is-multiline">
        <div class="column is-4">
          <div class="container">
            <figure class="image">
              <img src="../assets/svg/about.svg">
            </figure>
          </div>
        </div>
      <div class="column is-8">
      <div class="container has-text-left">
          <Story
            v-for="about in abouts"
            :key="about.id"
            v-bind:story="about.story"
            v-bind:title="about.title"
          />
      </div>
      <div class="container has-text-left">
        <h2 class="subtitle">
          We're a team of gifted and curious people building the future of customer success.
        </h2>
        <div class="tile is-ancestor">
          <Member
            v-for="member in members" 
            :key="member.id"
            v-bind:profile="member.profile"
            v-bind:username="member.username"
            v-bind:bio="member.bio"
            v-bind:avatar="member.avatar"
            v-bind:name="member.name"
            v-bind:title="member.title"
          />
        </div>
      </div>
      </div>

      </div>

    </div>

  </section>
</template>
<script>
import api from '../api';
const abouts = api.getAbout();

import Story from '../components/Story.vue';
import Member from '../components/Member.vue';

const members = api.getTeam();

export default {
  name: 'About',
  components: {
    Story,
    Member,
  },
  data() {
    return {
     abouts: abouts,
     members: members,
    };
  },
  mounted() {
  },
};
</script>
