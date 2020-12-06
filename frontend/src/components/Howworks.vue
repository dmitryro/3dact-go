<template>
      <div class="container is-fullwidth  has-text-left ml-2 mr-2 mt-5 mb-5">
      <fieldset class="fieldset">
            <legend> 
            <i class="has-text-info fas fa-circle fa-3x"><div style="color:white;margin-left:0.6em;margin-top:-1.4em;font-size:0.6em;">{{ index }}</div></i></legend>
             {{ description }}
            </fieldset>
      </div>
</template>

<script>
export default {
  name: 'Howworks',
  props: {
    index: String,
    title: String,
    description: String,
  },
};
</script>

<style>

</style>
