<script setup>
import { ref } from 'vue'
import { createCharacter } from '../libs/saveCharacter'

const cname = ref('')
const cimage = ref(null)
const selectedImage = ref(null)

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

const handleCreate = async () => {
  try {
    const result = await createCharacter(cname.value, cimage.value)
    alert(`キャラクター作成成功！ID: ${result.id}`)

  } catch (error) {
    alert(`エラー: ${error.message}`)
  }
  goToTrain()
}

const saveImage = (url) => {
  localStorage.setItem('c1', url)
  cimage.value = url
  selectedImage.value = url;
  console.log("Base64変換結果:", cimage.value.slice(0, 50) + "...")
}

</script>

<template>
  <div>
    <h1><span class="fri">とも</span><span class="life">ライフ</span></h1>
    <h2>育成するキャラを選択</h2>
    <div class="image-list">
      <!-- 画像1 -->
      <img 
        src="/niku.jpg" 
        alt="キャラ1"
        @click="saveImage('/niku.jpg')"
        :class="{selected: selectedImage === '/niku.jpg', dimmed:  selectedImage && selectedImage !== '/niku.jpg'}"
      />
      <!-- 画像2 -->
      <img 
        src="/mascot1.png" 
        alt="キャラ2"
        @click="saveImage('/mascot1.png')"
        :class="{selected: selectedImage === '/mascot1.png', dimmed:  selectedImage && selectedImage !== '/mascot1.png'}"
      />
      <!-- 画像3 -->
      <img 
        src="/penguin.png" 
        alt="キャラ3"
        @click="saveImage('/penguin.png')"
        :class="{selected: selectedImage === '/penguin.png', dimmed:  selectedImage && selectedImage !== '/penguin.png'}"
      />
    </div>
    <input v-model="cname" placeholder="キャラ名を入力"></input>
    <p></p>
    <button @click="handleCreate" class="btn">start</button>

  </div>
</template>

<style scoped>
.image-list {
  display: flex;
  gap: 15px;
}
img {
  width: 150px;
  height: auto;
  border: 2px solid transparent;
  border-radius: 12px;
  cursor: pointer;
  transition: border 0.2s;
}
img:hover {
  border: 2px solid #4caf50; /* hover時に枠をつける */
}
/* 選択された画像 */
.selected {
  border: 3px solid #4caf50;
  filter: brightness(1); /* 明るく */
  transform: scale(1.05);
}

.fri{
  color: green;
}

.life{
  color: orange;
}
/* 他の画像を暗くする */
.dimmed {
  filter: brightness(0.5);
}

.btn{
    color: white;
    display: inline-block;
    border: 2px solid white;
    background: red; 
    padding:o 8px;
    font-size:32px;
}

input{
    width:352px;
    height:32px;
    font-size:32px;
    background:white;
    color:black;
    border-radius: 20px;
    margin-top: 20px;
  }
</style>