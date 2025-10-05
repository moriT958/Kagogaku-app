<script setup>
  import { ref } from 'vue';
  // import Train from './components/Train.vue'
  import { createCharacter } from '../libs/saveCharacter';

  const cname=ref('');
  const cimage=ref(null);

  const goToTrain=()=>{
    window.location.hash='/Train'
  }

  const convertUrlToBase64 = async (url) => {
    const response = await fetch(url)
    const blob = await response.blob()
    const reader = new FileReader()
    return new Promise((resolve) => {
      reader.onloadend = () => resolve(reader.result)
      reader.readAsDataURL(blob)
  })
}
  
  const imageList=[
    '/mascot1.png',
    'https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjs8AvZeLXqN9PMHmzpxBpXq5033bksBQHA3PD35qpdKvH1Rs6E30OYoe0u4Bpch_7tv0cBy6jJLSOzCwDKDGUlrJybLhGyqlbWGe9wFE5_3i6ccR2C3TVj9Tq-rvd8P9CT7VT3aof5jBDP/s180-c/buranko_girl_sad.png',
    'https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhFEro64Nm55L1dVyHRAcfxmBk10XUp8LhnimtcACnq4y3NcQ2CL90khDUTdwXKhlAc4bSXDZ79Zuo8rLSazzCwH61-iuyQgTLCaDtb5Ol4kJfPvswy2qOa1lbO1-RiIA8rMf_VkuRdBnKQ/s180-c/sori_snow_boy.png'
  ]

  const selectedIndex=ref(null)

  const handleCreate = async () => {
  try {
    const result = await createCharacter(cname.value, cimage.value)
    alert(`キャラクター作成成功！ID: ${result.id}`)
    goToTrain()

  } catch (error) {
    alert(`エラー: ${error.message}`)
  }
}

  function handleClick(img,index){
    selectedIndex.value=index
    localStorage.setItem('c1', img)
    cimage.value = img;
    // alert(`選択した画像を保存しました: ${index}`);
  }
</script>

<template>
  <p>好きなキャラを選んでね！</p>
  <div class="images">
    <img
      v-for="(img,index) in imageList"
      :key="index"
      :src="img"
      :class="{selected: selectedIndex === index}"
      @click="handleClick(img,index)"
  </div>
  <p>キャラ名</p>
  <input v-model="cname" placeholder="キャラ名を入力"></input></br>
  <button @click="handleCreate" class="start">Start</button>
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
    width:352px;
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
    border: 4px solid #4caf50;
  }
</style>