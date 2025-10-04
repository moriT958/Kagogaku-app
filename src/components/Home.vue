<script setup>
  import {ref} from 'vue';

  const textInput=ref('');
</script>

<template>
  <p>好きなキャラを選んでね！</p>
  <div class="images">
    <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
    <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
    <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
    <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
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