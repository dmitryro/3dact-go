<template>
<swiper-slide>
  <div class="column is-fullwidth">
    <div v-tilt class="box">
      <div class="card-content">
        <div class="media card-adjusted">
          <div class="media-content  has-text-centered">
            <p class="title">{{ title }}</p>
          </div>
        </div>
        <div class="content has-text-left">
          <p id="desc">{{ description }}</p>
          <ul>
            <li v-for="point in points" :key="point.id">
                {{ point }}
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</swiper-slide>
</template>

<script>
//import Bulleted from '../components/Bulleted.vue';
import { SwiperSlide } from 'vue-awesome-swiper'
export default {
  name: 'Article',
  components: {
    SwiperSlide,
    //Bulleted,
  },
  props: {
    title:String,
    description:String,
    points:Array
  },
  data() {
    return {
    };
  },
  mounted() {
  }
};
</script>

<style scoped>
  .box {
    transform-style: preserve-3d;
    transform: perspective(1000px);
  }

  .media {
    transform: translateZ(20px);
  }

  .stars-counter {
    color: #f1c40f
  }

  #desc {
    padding-top: 30px;
    word-wrap: break-word;
    overflow-wrap: break-word;
  }

  .card-footer {
    margin-top: 10px;
    transform: translateZ(20px);
  }

  .card-footer a {
    color: hsl(0, 0%, 21%);
  }

</style>
