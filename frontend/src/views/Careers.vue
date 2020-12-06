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
      <div class="container has-text-left">
        <h1 class="title">
          Careers
        </h1>
        <h2 class="subtitle">
          At the present time we're not hiring yet.
        </h2>
        <div class="tile is-ancestor">
        </div>
      </div>
      </div>

      </div>

    </div>
  </section>
</template>

<script>
import api from '../api';
//import Member from '../components/Member.vue';

export default {
  name: 'About',
  components: {
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
