<template>
      <div class="container is-fullwidth  has-text-left ml-2 mr-2 mt-5 mb-5">

      <fieldset class="fieldset">
            <legend>
            <i class="has-text-info fas fa-circle fa-3x"><div style="color:white;margin-left:0.6em;margin-top:-1.4em;font-size:0.6em;">{{ index }}</div></i></legend>
             <div><span class="icon is-small has-text-danger local">Q</span>: {{ question }}</div>
             <div><span class="icon is-small has-text-success local">A</span>: {{ answer }} </div>
      </fieldset>
      </div>
</template>

<script>
export default {
  name: 'Faqs',
  props: {
    index: String,
    question: String,
    answer: String,
  },
};
</script>

<style>
  .local {
     font-weight:bold;
     font-size:1.2em;
   }
  .top {
     color: #CCCCCC;
     border: 1px solid #CDCDCD; 
     border-radius: 5px;
  }
</style>
~                      
