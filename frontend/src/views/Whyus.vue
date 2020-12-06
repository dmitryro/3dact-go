<template>
  <section class="hero">
    <div class="hero-body">
      <div class="container has-text-left">
          <div class="columns">
              <div class="column is-1"></div>
              <div class="column is-11 has-text-left">
                  <h1 class="title">Why Us</h1>
                  <p class="subtitle">Here's our case for our technology</p>
              </div>
          </div>
          <div class="columns">
                <div class="column is-1"></div>
                <div class="column is-5">
                   <div class="container">
                       <figure class="image is-fullwidth">
                           <img src="../assets/svg/whyus.svg">
                       </figure>
                   </div>
                   <div class="container mt-6">
                   </div>                                
                </div> <!-- end column-->
                <div class="column is-5">
                   <Whyuseus
                    v-for="why in whyus"
                    :key="why.id"
                    v-bind:index="why.index"
                    v-bind:title="why.title"
                    v-bind:list="why.list"
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
import Whyuseus from '../components/Whyuseus.vue';
const whyus = api.getWhyus();

export default {
  name: 'Whyus',
  components: {
     Whyuseus,
  },
  data() {
    return {
      whyus: whyus,
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
  .whydisc {
     color:#1B7DF0;
     font-size:0.8em;
   }
  .whyblock {
     margin-bottom: 1.0em; 
   }

  .whylist {
    list-style: disc;
    font-size: 20px;
    padding-left: 1.0em;
    color: #1B7DF0;
  }
</style>
