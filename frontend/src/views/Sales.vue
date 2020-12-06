<template>
  <section class="hero">
    <div class="hero-body">
      <div class="container has-text-left">
          <div class="columns">
              <div class="column is-1"></div>
              <div class="column is-11 has-text-left">
                  <h1 class="title">Sales</h1>
                  <p class="subtitle">Some of the most frequently asked questions answered</p>
              </div>
          </div>
          <div class="columns">
                <div class="column is-1"></div>
                <div class="column is-5">
                   <div class="container">
                       <figure class="image is-fullwidth">
                           <img src="../assets/svg/sales.svg">
                       </figure>
                   </div>
                   <div class="container mt-6">
                   </div>                                
                </div> <!-- end column-->
                <div class="column is-5">

                 <span class="title">Advanced knowledge management for productive sales teams.</span>

                 <ul class="bulletless">
                 <li>Give future customers and current sales teams members the tools they need</li>
                 <li>AI-driven knowledge management shortens sales cycles, boosts productivity.</li>
                 <li>Automated workflows optimize time usage and allow a more human interaction.</li>
                 </ul>
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
