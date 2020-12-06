<template>
  <section class="hero">
    <div class="hero-body">
      <div class="container has-text-left">
          <div class="columns">
              <div class="column is-1"></div>
              <div class="column is-11 has-text-left">
                  <h1 class="title">Frequently Asked Questions</h1>
                  <p class="subtitle">Some of the most frequently asked questions answered</p>
              </div>
          </div>
          <div class="columns">
                <div class="column is-1"></div>
                <div class="column is-5">
                   <div class="container">
                       <figure class="image is-fullwidth">
                           <img src="../assets/svg/questions.svg">
                       </figure>
                   </div>
                   <div class="container mt-6">


                   </div>                                
                </div> <!-- end column-->
                <div class="column is-5">
                   <Faqs
                    v-for="faq in faqs"
                    :key="faq.id"
                    v-bind:index="faq.index"
                    v-bind:question="faq.question"
                    v-bind:answer="faq.answer"
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
const faqs = api.getFaqs();

import Faqs from '../components/Faqs.vue';




export default {
  name: 'Project',
  components: {
      Faqs,
  },
  data() {
    return {
      faqs: faqs,
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
