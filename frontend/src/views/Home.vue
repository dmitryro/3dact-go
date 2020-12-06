<template>

  <section class="hero is-fullheight">
    <div class="hero-body">
         <div class="columns">
            <div class="column is-5">
              <div class="container">
                <figure class="image is-fullwidth">
                  <img src="../assets/svg/home.svg">
                </figure>
              </div>
            </div>
            <div class="column is-7">
               <div class="container has-text-centered">

               <swiper ref="mySwiper" 
                       :auto-update="true"
                       :auto-destroy="true"
                       :delete-instance-on-destroy="true"
                       :cleanup-styles-on-destroy="true"
                       @ready="handleSwiperReadied"
                       @click-slide="handleClickSlide"
                       v-model="index"
                       :options="swiperOptions">
                          <Article
                             v-for="article in articles"
                             :key="article.id"
                             v-bind:points="article.points"
                             v-bind:title="article.title"
                             v-bind:description="article.description"
                          />
                         <div class="swiper-pagination swiper-pagination-blue" slot="pagination"></div>
               </swiper>

               </div> 
            </div> 
          </div> 
      </div>











    <div class="hero-body">
         <div class="columns">

            <div class="column is-6">
             
              <div class="container">


            <svg class="slack" viewBox="0 0 842 596">
                <path fill="#fff" d="M172 0 h670 v595.3 H172 z"/>
                <path fill="#4d394b" d="M0 0 h172 v595.3 H0 z" class="sidebar"/>
                <g class="title-section">
                    <path fill="#fff" fill-opacity="0.4"
                        d="M22.3 23 h127.4 v21 H22.3 z
                        M50.6 61.2 H122 v3 H50.7 z"/>
                    <ellipse cx="35" cy="62.8" rx="6.7" ry="6.8" fill="#52b98c"/>
                </g>
                <path class="channelTitle" fill="#fff" fill-opacity=".4" d="M22.7 112 H150 v9 H22.8 z"/>
                <g class="channelBars" fill="#fff" fill-opacity=".4">
                    <path class="cb1" d="M30.4 137.6 H102 v3 H30.3 z" />
                    <path class="cb2" d="M30.4 164.2 H102 v3 H30.3 z" />
                    <path class="cb3" d="M30.4 190.8 H102 v3 H30.3 z" />
                    <path class="cb4" d="M30.4 215.0 H102 v3 H30.3 z" />
                    <path class="cb5" d="M30.3 240.8 h71.5 v3 H30.3 z" />
                    <path class="cb6" d="M30.3 267.8 h71.5 v3 H30.3 z" />
                    <path class="cb7" d="M30.3 294.4 h71.5 v3 H30.3 z" />
                    <path class="cb8" d="M30.3 321 h71.5 v3 H30.3 z" />
                </g>
                <path class="activeChannel" d="M132 215.2 c0-1.8-1.3-3.2-3-3.2 H0 v9 h129 c1.7 0 3-1.4 3-3.2 v-2.6 z" fill="#52b98c"/>
                <path class="directMessageTitle" fill="#fff" fill-opacity=".4" d="M22.3 366.5 h127.4 v9 H22.3 z"/>
                <g class="messageBars" fill="#fff" fill-opacity=".4">
                    <path d="M52.7 399 h71.5 v3 H52.7 z
                        M52.7 425.5 h71.5 v3 H52.7 z
                        M52.7 452.2 h71.5 v3 H52.7 z
                        M52.7 533.6 h71.5 v3 H52.7 z
                        M52.7 478.8 h71.5 v3 H52.7 z
                        M52.7 560.2 h71.5 v3 H52.7 z
                        M52.7 505.4 h71.5 v3 H52.7 z"/>
                </g>
                <g class="directMessage" fill="#52b98c">
                    <ellipse class="dm1" cx="37.1" cy="400.4" rx="6.7" ry="6.8" />
                    <ellipse class="dm2" cx="37.1" cy="427.1" rx="6.7" ry="6.8" />
                    <ellipse class="dm3" cx="37.1" cy="453.7" rx="6.7" ry="6.8" />
                    <ellipse class="dm4" cx="37.1" cy="480.4" rx="6.7" ry="6.8" />
                    <ellipse class="dm5" cx="37.1" cy="507" rx="6.7" ry="6.8" />
                    <ellipse class="dm6" cx="37.1" cy="535.2" rx="6.7" ry="6.8" />
                    <ellipse class="dm7" cx="37.1" cy="561.8" rx="6.7" ry="6.8" />
                </g>
                <path class="Avatar-1" d="M288.7 56 c0-5.3-4.2-9.5-9.4-9.5 h-44.6 c-5 0-9.3 4.2-9.3 9.4 v39 c0 5 4.2 9.3 9.3 9.3 h44.6 c5.2 0 9.4-4.2 9.4-9.4 V56 z" fill="#ea5a3d"/>
                <path class="Avatar-2" d="M288.7 56 c0-5.3-4.2-9.5-9.4-9.5 h-44.6 c-5 0-9.3 4.2-9.3 9.4 v39 c0 5 4.2 9.3 9.3 9.3 h44.6 c5.2 0 9.4-4.2 9.4-9.4 V56 z" fill="#785eb0"/>
                <path class="Avatar-3" d="M288.7 56 c0-5.3-4.2-9.5-9.4-9.5 h-44.6 c-5 0-9.3 4.2-9.3 9.4 v39 c0 5 4.2 9.3 9.3 9.3 h44.6 c5.2 0 9.4-4.2 9.4-9.4 V56 z" fill="#00b7e5"/>
                <path class="Avatar-4" d="M288.7 56 c0-5.3-4.2-9.5-9.4-9.5 h-44.6 c-5 0-9.3 4.2-9.3 9.4 v39 c0 5 4.2 9.3 9.3 9.3 h44.6 c5.2 0 9.4-4.2 9.4-9.4 V56 z" fill="#ffc10e"/>
                <path class="Avatar-5" d="M288.7 56 c0-5.3-4.2-9.5-9.4-9.5 h-44.6 c-5 0-9.3 4.2-9.3 9.4 v39 c0 5 4.2 9.3 9.3 9.3 h44.6 c5.2 0 9.4-4.2 9.4-9.4 V56 z" fill="#c67bd1"/>

                <g class="Block-1" fill-opacity=".3">
                    <path class="top-1" d="M303 59h263.6v9H303z"/>
                    <path class="bottom-1" d="M303 82.7h392.6v9H303z"/>
                </g>
                <g class="Block-2" fill-opacity=".3">
                    <path class="top-2" d="M303 59h263.6v9H303z"/>
                    <path class="bottom-2" d="M303 82.7h392.6v9H303z"/>
                </g>
                <g class="Block-3" fill-opacity=".3">
                    <path class="top-3" d="M303 59h263.6v9H303z"/>
                    <path class="bottom-3" d="M303 82.7h392.6v9H303z"/>
                </g>
                <g class="Block-4" fill-opacity=".3">
                    <path class="top-4" d="M303 59h263.6v9H303z"/>
                    <path class="bottom-4" d="M303 82.7h392.6v9H303z"/>
                </g><g class="Block-5" fill-opacity=".3">
                    <path class="top-5" d="M303 59h263.6v9H303z"/>
                    <path class="bottom-5" d="M303 82.7h392.6v9H303z"/>
                </g>
            </svg>






              </div>
            </div>
            <div class="column is-6">

                  <h1 class="title">Slack-first productivity application</h1>
                   <p class="subtitle">Your knowledge your way.</p>
                   <br/>
                   <p> 3D Act is an AI-driven knowledge management platform that allows you
                       deliver your knowledge to every corner of the endless online world.</p>

            </div>
         </div>
    </div>
  </section>

</template>
<script>
import '../assets/css/swiper.css'
import api from '../api';
import Article from '../components/Article.vue';
const articles = api.getFrontArticles();
import Vue from 'vue'
import { Swiper as SwiperClass, Pagination, Mousewheel, Autoplay } from 'swiper/core'
import getAwesomeSwiper from 'vue-awesome-swiper/dist/exporter'
// Swiper modules
SwiperClass.use([Pagination, Mousewheel, Autoplay])
// -------------------------------------------------
// Global use
Vue.use(getAwesomeSwiper(SwiperClass))
import { directive } from 'vue-awesome-swiper'
const { Swiper } = getAwesomeSwiper(SwiperClass)
// `@click-slide` event has special treatment for Swiper's loop mode, which is still available in loop mode
export default {
  name: 'Home',
  components: {
    Article,
    Swiper,
  },
  data() {
    return {
      articles: articles,
      index: 0,
      swiperOptions: {
                   notNextTick: true,
                    loop: true,
                    initialSlide: 0,
                    autoplay: {
                        delay: 13500,
                        disableOnInteraction: true
                    },
                    speed: 1000,
                    grabCursor: true,
                    pagination: {


            el: '.swiper-pagination',
            clickable: true,
            //renderBullet(index, className) {
            //  return `<span class="${className} swiper-pagination-blue">${index + 1}</span>`
           // }


            //           el: ".swiper-pagination",
            //            clickable: true,
            //            type: "bullets"
                    }
      }
    };
  },
  methods: {
      navigate() {
      },
      handleSwiperReadied(swiper) {
         console.log('Swiper was munted!', swiper)
      },
      handleClickSlide(index, reallyIndex) {
         this.index = index;
         console.log('Click slide!', index, reallyIndex)
      }
  },
  computed: {
      swiper() {
        return this.$refs.mySwiper.$swiper
      }
  },
  directives: {
    swiper: directive
  },
  mounted() {
    console.log('Current Swiper instance object', this.swiper)
    //this.swiper.slideTo(3, 1000, false)
  }
};
</script>
<style>
  footer {
    margin-top: 50px;
  }
</style>
