<template>
  <section class="hero">
    <div class="hero-body">
      <div class="container has-text-left">
          <div class="columns">
              <div class="column is-1"></div>
              <div class="column is-11 has-text-left">
                  <h1 class="title">How it all works</h1>
                  <p class="subtitle">See how the parts of the puzzle fit together into a perfect symphony</p> 
              </div>
          </div>
          <div class="columns">
                <div class="column is-1"></div>
                <div class="column is-5">
                   <div class="container">
                       <figure class="image is-fullwidth">
                           <img src="../assets/svg/howitworks.svg">
                       </figure>
                   </div>
                </div> <!-- end column-->
                <div class="column is-5">

          <Howworks
            v-for="how in howworks"
            :key="how.id"
            v-bind:index="how.index"
            v-bind:title="how.title"
            v-bind:description="how.description"
          />

                </div>
                <div class="column is-1"></div>
             </div> <!-- end columns -->
   
          </div><!-- end container -->
      </div>
  </section>
</template>

<script>

import api from '../api';
const howworks = api.getHowworks();

import Howworks from '../components/Howworks.vue';

export default {
  name: 'Howitworks',
  components: {
     Howworks,
  },
  data() {
    return {
      howworks: howworks,
    };
  },
  mounted() {
  },
};
</script>

<style>
  footer {
    margin-top: 50px;
  }
</style>
