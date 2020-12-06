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
        <h1 class="title">
          Team
        </h1>
        <h2 class="subtitle">
          We're a team of highly skilled and experienced individuals that builds the future.
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
const members = api.getTeam();

import Member from '../components/Member.vue';

export default {
  name: 'About',
  components: {
    Member,
  },
  data() {
    return {
      members: members,
    };
  },
  mounted() {
  },
};
</script>
