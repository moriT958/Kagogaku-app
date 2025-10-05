<script setup>
import { computed, onMounted, ref } from 'vue'

// 時刻を表示するための変数
const sleepTime = ref(null)
const wakeTime = ref(null)
const sleepingTime = ref(null)
const selectedImage = ref(null)
const characterName = ref('')


onMounted(() => {
  const saved = localStorage.getItem("c1")
  if (saved) {
    const base = import.meta.env.BASE_URL // "/Kagogaku-app/"
    selectedImage.value = base + saved.replace(/^\//, '')
    console.log("Resolved image path:", selectedImage.value)
  }
})

// ボタンクリック時の関数
const recordSleep = () => {
  const now = new Date()
  sleepTime.value = now.toLocaleTimeString() // 現在時刻を文字列で保存
}

const recordWake = () => {
  const now = new Date()
  wakeTime.value = now.toLocaleTimeString()
}

const number = ref(0)
const cname = ref('')

// 数値に応じて文字列を返す
const message = computed(() => {
  if (number.value < 30) {
    return '悪い'
  } else if (number.value < 60) {
    return '普通'
  } else {
    return '良い'
  }
})
</script>

<template>
  <div class="all">
    <input v-model="cname"></input>
    <div class="chara-status">
      <span class="cname">キャラ名: {{ cname }}</span>
      <span class="cstatus">健康状態: {{ message }}</span>
    </div>
    
    <div v-if="selectedImage" class="chara_img">
      <img :src="selectedImage" alt="キャラ画像" width="200" />
    </div>
    <div v-else>
      <p>キャラがまだ選択されていません。</p>
    </div>
    
    <div class="buttons">
      <button @click="recordSleep" class="sleepbtn">
        <img src="/sleep_cat.png" class="btn_img"></img>
        <p class="btn_t">就寝</p>
      </button>
      <button @click="recordWake" class="wakebtn">
        <img src="/wake_cat.png" class="btn_img"></img>
        <p class="btn_t">起床</p>
      </button>
      <!-- <input id="meal"></input> -->
      <input id="test" type="number" v-model.number="number" class="tinput">
    </div>

    <div class="status">
      <p v-if="sleepTime" class="time">就寝時刻: {{ sleepTime }}</p>
      <p v-if="wakeTime" class="time">起床時刻: {{ wakeTime }}</p>
    </div>

  </div>
</template>

<style scoped>
.all{
  width: 100%;
  height: auto;
  background-color: beige;
}

.chara-status{
  display: flex;
  justify-content: space-between;
}

.cname{
  /* border: 1px solid green; */
  background-color: pink;
  border-radius: 20px;
  padding: 5px;
}

.cstatus{
  border: 1px solid blue;
  border-radius: 20px;
  padding: 5px;
}

img{
  width: 200px;
  height: auto;
}

.btn_t{
  font-size: 20px;
}

.btn_img{
  width: 40px;
  height: auto;
}

.chara_img{
  margin-bottom: 50px;
  /*background-color: black;*/
  width: max;
  height: auto;
}

.buttons{
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: space-between;
  background-color: lightblue;
  border-radius: 12px;
}

.sleepbtn{
  color: blue;
  margin: 10px;
}

.wakebtn{
  color: red;
  margin: 10px;
}

.tinput{
  margin: 20px;
  border-radius: 20px;

}

.status{
  border: 2px solid yellowgreen;
  border-radius: 20px;
  background-color: white;
}

.time{
  color: yellowgreen;
}
</style>
