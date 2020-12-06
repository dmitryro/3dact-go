<template>
      <div class="container multiline has-text-left ml-2 mr-2">
        <fieldset class="fieldset">
         <legend>
            {{ name }}
         </legend>
            <figure class="image container is-180x180">
              <img class="is-rounded" src="../assets/png/icon-avatar-default.png">
            </figure>
            <div class="container has-text-centered">
            <span class="smaller">
            {{ title }}
            </span>
            </div>
            <hr class="navbar-divider" style="z-index: 100;">
            <div class="container" style="font-size:0.92em;">
            {{ bio }}
            </div>
         </fieldset>
      </div>
</template>

<script>
export default {
  name: 'Member',
  props: {
    title: String,
    name: String,
    profile: String,
    username: String,
    avatar: String,
    bio: String,
  },
};
</script>

<style>

</style>
