<template>
  <div class="column is-one-third">
    <div v-tilt class="box">
      <div class="card-content">
        <div class="media card-adjusted">
          <div class="media-content  has-text-centered">
            <p class="title is-4">{{ title }}</p>
          </div>
        </div>
        <div class="content">
          <p id="desc">{{ description }}</p>
          <p class="subtitle is-6" style="font-size:0.8em;margin-top:-0.8em">{{ packagenote }}</p>


          <nav class="level is-mobile">

            <ul class="has-text-left" style="  list-style: none;  list-style-type: none; font-size:0.95em;">
            <div class="columns is-multiline is-mobile" v-for="item in items" :key="item">
                 <li>
                         <div class="is-full media-content  has-text-left">
                          <span class="icon is-small has-text-primary">
                                <i class="fas fa-check"></i>
                          </span>
                          <span class="ml-3">{{ item }}</span>
                          </div>
                 </li>
            </div>
            </ul>
          </nav>

          <nav class="level is-mobile">
            <div class="level-item has-text-left">

            <div class="level-item has-text-centered">
              <div>
                <p class="heading has-text-left">Price</p>
                <p class="title is-4">{{ price }}</p>
              </div>
            </div>
            <div class="level-item has-text-centered">
              <div>
                <p class="heading has-text-left">Period</p>
                <p class="title is-4 stars-counter">{{ period }}</p>
              </div>
            </div>
            <div class="level-item has-text-centered">
              <div>
                <p class="heading has-text-left">Team</p>
                <p class="title is-4 stars-counter">Any</p>
              </div>
            </div>


            </div>

          </nav>

         
          <div class="buttons is-centered">
                       <b-tooltip
                        type="is-light"
                        label="Try it before you buy it"
                        position="is-bottom"
                        size="is-medium"
                        multilined
                       >
                       <b-button
                        tag="router-link" :to="{ name: 'pricing' }"
                        type="is-primary"
                        class="is-rounded"
                        size="is-medium"
                       >
                         Get started
                      </b-button>
                     </b-tooltip>
         </div>

        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Package',
  props: {
    title: String,
    description: String,
    price: String,
    period: String,
    pricenote: String,
    packagenote: String,
    items:Array,
  },
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
