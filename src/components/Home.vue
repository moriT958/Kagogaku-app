<script setup>
  import { ref } from 'vue';
  // import Train from './components/Train.vue'
  const textInput=ref('');
  const goToTrain=()=>{
    window.location.hash='/Train'
  }
  const saveImage = (url) => {
    localStorage.setItem('c1', url)
    // alert(`選択した画像を保存しました: ${url}`);
  }
  // const saveName = (cName) => {
  //   localStorage.setItem('cn', cName)
  // }
  const images=document.querySelectorAll('.images img');

  images.forEach(img=>{
    img.addEventListener('click',()=>{
      images.forEach(i=>i.classList.remove('selected'));
      img.classList.add('selected');
    });
  });
</script>

<template>
  <p>好きなキャラを選んでね！</p>
  <div class="images">
    <!-- 画像1 -->
    <img 
        src="https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEi_KDuCUJ4SXPSs6ZVz2LFbkWyS4XTccFEfiizRuoUPM9PmR1Ob9P7IJpFG4hb2_wloVK6eMLvN8bwv8ILauUQWacxn7DzeHCvuQudsNZucMUTrzaMdHRzwjBJ5pDVJNTMhwAnDegyKBtAs6wEjFgsc2HY-9q4WoCHVuBzYIQ_Xv1KPBK8ZBHyx9EpGX3PA/s180-c/bucket_boy.png" 
        alt="キャラ1"
        @click="saveImage('https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEi_KDuCUJ4SXPSs6ZVz2LFbkWyS4XTccFEfiizRuoUPM9PmR1Ob9P7IJpFG4hb2_wloVK6eMLvN8bwv8ILauUQWacxn7DzeHCvuQudsNZucMUTrzaMdHRzwjBJ5pDVJNTMhwAnDegyKBtAs6wEjFgsc2HY-9q4WoCHVuBzYIQ_Xv1KPBK8ZBHyx9EpGX3PA/s180-c/bucket_boy.png')"
      />
      <!-- 画像2 -->
      <img 
        src="https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjs8AvZeLXqN9PMHmzpxBpXq5033bksBQHA3PD35qpdKvH1Rs6E30OYoe0u4Bpch_7tv0cBy6jJLSOzCwDKDGUlrJybLhGyqlbWGe9wFE5_3i6ccR2C3TVj9Tq-rvd8P9CT7VT3aof5jBDP/s180-c/buranko_girl_sad.png" 
        alt="キャラ2"
        @click="saveImage('https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjs8AvZeLXqN9PMHmzpxBpXq5033bksBQHA3PD35qpdKvH1Rs6E30OYoe0u4Bpch_7tv0cBy6jJLSOzCwDKDGUlrJybLhGyqlbWGe9wFE5_3i6ccR2C3TVj9Tq-rvd8P9CT7VT3aof5jBDP/s180-c/buranko_girl_sad.png')"
      />
      <!-- 画像3 -->
      <img 
        src="https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhFEro64Nm55L1dVyHRAcfxmBk10XUp8LhnimtcACnq4y3NcQ2CL90khDUTdwXKhlAc4bSXDZ79Zuo8rLSazzCwH61-iuyQgTLCaDtb5Ol4kJfPvswy2qOa1lbO1-RiIA8rMf_VkuRdBnKQ/s180-c/sori_snow_boy.png" 
        alt="キャラ3"
        @click="saveImage('https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhFEro64Nm55L1dVyHRAcfxmBk10XUp8LhnimtcACnq4y3NcQ2CL90khDUTdwXKhlAc4bSXDZ79Zuo8rLSazzCwH61-iuyQgTLCaDtb5Ol4kJfPvswy2qOa1lbO1-RiIA8rMf_VkuRdBnKQ/s180-c/sori_snow_boy.png')"
      />
  </div>
  <p>キャラ名</p>
  <input type="text" v-model="textInput" />
  <p></p>
  <button @click="goToTrain" class="start">はじめる</button>
</template>

<style>
  p{
    font-size:32px;
  }

  img{
    margin:8px;
    background:white;
  }

  .darkened{
    filter:brightness(50%);
  }
  
  input{
    width:320px;
    height:32px;
    font-size:32px;
    background:white;
    color:black;
  }

  .start{
    display: inline-block;
    border: 2px solid white;
    background: red; 
    padding:o 8px;
    font-size:32px;
  }

  .start:hover{
    filter:brightness(80%);
    border:2px solid rgba(255,255,255,0.6);
  }

  img:hover {
    border: 4px solid #4caf50; /* hover時に枠をつける */
  }

  .selected{
    filter:brightness(50%);
  }
</style>