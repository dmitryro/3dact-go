<template>
  <div class="column is-one-third">
    <div v-tilt class="box">
      <div class="card-content">
        <div class="media">
          <div class="media-left">
            <figure class="image is-64x64">
              <img :src="avatar" alt="User avatar">
            </figure>
          </div>
          <div class="media-content">
            <p class="title is-4">{{ repository }}</p>
            <p class="subtitle is-6">@{{ profile }}</p>
          </div>
        </div>
        <div class="content has-text-centered">
          <p id="desc">{{ description }}</p>
          <nav class="level is-mobile">
            <div class="level-item has-text-centered">
              <div>
                <p class="heading">Price</p>
                <p class="title is-4">{{ lang }}</p>
              </div>
            </div>
            <div class="level-item has-text-centered">
              <div>
                <p class="heading">Per</p>
                <p class="title is-4 stars-counter">{{ stars }}</p>
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
  name: 'Repository',
  props: {
    avatar: String,
    repository: String,
    profile: String,
    description: String,
    lang: String,
    stars: Number,
    forks: Number,
    url: String,
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
