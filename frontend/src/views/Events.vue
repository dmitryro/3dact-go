<template>
  <section class="hero is-fullheight">
    <div class="hero-body">
      <div class="columns is-multiline">
        <div class="column">
          <div class="container">
            <figure class="image">
              <img src="../assets/svg/about.svg">
            </figure>
          </div>
        </div>

      <div class="column">
      <div class="container has-text-centered">
        <h1 class="title">
          Events
        </h1>
        <h2 class="subtitle">
          Nous sommes une communauté de personnes intéressées par l'intelligence artificielle
        </h2>
        <div class="tile is-ancestor">
          <Member
            v-for="member in members" 
            :key="member.id"
            v-bind:profile="member.profile"
            v-bind:username="member.username"
            v-bind:bio="member.bio"
            v-bind:avatar="member.avatar"
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
import Member from '../components/Member.vue';

export default {
  name: 'About',
  components: {
    Member,
  },
  data() {
    return {
      members: api.getMembers(),
    };
  },
  mounted() {
  },
};
</script>
