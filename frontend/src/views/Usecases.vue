<template>
  <section class="hero">
    <div class="hero-body">
      <div class="container has-text-left">
          <div class="columns">
              <div class="column is-1"></div>
              <div class="column is-11 has-text-left">
                  <h1 class="title">Use Cases</h1>
                  <p class="subtitle">Some most frequent use cases</p>
              </div>
          </div>
          <div class="columns">
                <div class="column is-1"></div>
                <div class="column is-5">
                   <div class="container">
                       <figure class="image is-fullwidth">
                           <img src="../assets/svg/usecases.svg">
                       </figure>
                   </div>
                   <div class="container mt-6">
                   </div>                                
                </div> <!-- end column-->
                <div class="column is-5">
                   <Usecase
                    v-for="usecase in usecases"
                    :key="usecase.id"
                    v-bind:index="usecase.index"
                    v-bind:title="usecase.title"
                    v-bind:description="usecase.description"
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
import Usecase from '../components/Usecase.vue';
const usecases = api.getUsecases();

export default {
  name: 'Usecases',
  components: {
     Usecase,
  },
  data() {
    return {
      usecases: usecases,
    };
  },
  mounted() {
  },

};
</script>
<style>
  .footer {
    margin-top: 50px;
  }

  .listitem {
   }

  .whybullet {
    color: #707070;
  }

  .left {
     float:left;
     width:5%;
   }

  .right {
     float:left;
     width:94%;
   }

   .clear {
     clear:both;
   }
  .choosedisc {
     color:#1B7DF0;
     font-size:0.8em;
   }
  .chooseblock {
     margin-bottom: 1.0em; 
   }
  .chooselist {
    list-style: disc;
    font-size: 20px;
    padding-left: 1.0em;
    color: #1B7DF0;
  }
  .legendheader {
    font-weight:bold;
    font-size:1.4em;
    margin-bottom: 1.0em;
  }
  .adjust {
    margin-top: -1.7em;
  }
</style>
