<script setup>
  import { ref } from 'vue';

  const textInput=ref('');
</script>

<template>
  <p>好きなキャラを選んでね！</p>
  <div class="images">
    <img src="https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEi_KDuCUJ4SXPSs6ZVz2LFbkWyS4XTccFEfiizRuoUPM9PmR1Ob9P7IJpFG4hb2_wloVK6eMLvN8bwv8ILauUQWacxn7DzeHCvuQudsNZucMUTrzaMdHRzwjBJ5pDVJNTMhwAnDegyKBtAs6wEjFgsc2HY-9q4WoCHVuBzYIQ_Xv1KPBK8ZBHyx9EpGX3PA/s180-c/bucket_boy.png" class="logo vue" alt="Vue logo" />
  </div>
    <p>キャラ名</p>
  <input type="text" v-model="textInput" />
  <div class="button" @click="click">はじめる</div>
</template>

<style>
  input{
    width:320px;
    height:32px;
  }

  .button{
    display: inline-block;
    border: 2px solid rgb(255,255,255);
    background: linear-gradient(to bottom, rgb(255,244,92), rgb(240,229,86)); 
  }
</style>