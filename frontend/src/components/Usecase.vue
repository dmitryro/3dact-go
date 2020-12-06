<template>
      <fieldset class="fieldset chooseblock">
            <legend><i class="has-text-info fas fa-circle fa-3x"><div style="color:white;margin-left:0.6em;margin-top:-1.4em;font-size:0.6em;">{{ index }}</div></i>
            </legend>
                   <div class="container legendheader has-text-weight-bold">
                         {{ title }}
                   </div>
                   <div class="columns is-fullwidth">
                        {{ description }}
                  </div>
     </fieldset>
</template>

<script>
export default {
  name: 'Usecase',
  props: {
    index: String,
    title: String,
    description: String
  },
};
</script>

<style>

</style>
