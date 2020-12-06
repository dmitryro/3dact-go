<template>
  <section class="hero">
    <div class="hero-body">
      <div class="container has-text-left">
          <div class="columns">
              <div class="column is-1"></div>
              <div class="column is-11 has-text-left">
                  <h1 class="title">Teams Solution</h1>
                  <p class="subtitle">Teams</p>
              </div>
          </div>
          <div class="columns">
                <div class="column is-1"></div>
                <div class="column is-5">
                   <div class="container">
                       <figure class="image is-fullwidth">
                           <img src="../assets/svg/teams.svg">
                       </figure>
                   </div>
                   <div class="container mt-6">
                   </div>                                
                </div> <!-- end column-->
                <div class="column is-5">
                 FAQ comes here
                </div>
                <div class="column is-1"></div>
             </div> <!-- end columns -->
   
          </div><!-- end container -->
      </div>
  </section>
</template>

<script>
//import api from '../api';


export default {
  name: 'Project',
  components: {
  },
  data() {
    return {
      infos: [],
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
