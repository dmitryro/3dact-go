<template>
      <fieldset class="fieldset chooseblock">
            <legend><i class="has-text-info fas fa-circle fa-3x"><div style="color:white;margin-left:0.6em;margin-top:-1.4em;font-size:0.6em;">{{ index }}</div></i>
            </legend>
                   <div class="container legendheader has-text-weight-bold">
                         {{ title }}
                   </div>
                   <div class="columns is-fullwidth">
                        <ul>
                             <li v-for="item in list" :key="item.id" style="margin-top:0.5em;margin-bottom:0.5em;">                   
                               <div class="columns is-fullwidth">
                                     <div class="column is-1" ><span class="fas fa-circle is-small is-primary ml-4 choosedisc"></span></div>
                                     <div class="column is-11"><span>{{ item }}</span></div>
                               </div>
                             </li>
                        </ul>
                  </div>
     </fieldset>
</template>

<script>
export default {
  name: 'Howchoose',
  props: {
    index: String,
    title: String,
    list: Array,
  },
};
</script>

<style>

</style>
