<script setup>


const saveImage = (url) => {
  localStorage.setItem('c1', url)
  alert(`選択した画像を保存しました: ${url}`)
}

</script>

<template>
  <div>
    <h2>キャラ画像を選択</h2>
    <div class="image-list">
      <!-- 画像1 -->
      <img 
        src="/bucket_boy.png" 
        alt="キャラ1"
        @click="saveImage('/bucket_boy.png')"
      />
      <!-- 画像2 -->
      <img 
        src="/mascot1.png" 
        alt="キャラ2"
        @click="saveImage('/mascot1.png')"
      />
      <!-- 画像3 -->
      <img 
        src="https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhFEro64Nm55L1dVyHRAcfxmBk10XUp8LhnimtcACnq4y3NcQ2CL90khDUTdwXKhlAc4bSXDZ79Zuo8rLSazzCwH61-iuyQgTLCaDtb5Ol4kJfPvswy2qOa1lbO1-RiIA8rMf_VkuRdBnKQ/s180-c/sori_snow_boy.png" 
        alt="キャラ3"
        @click="saveImage('https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhFEro64Nm55L1dVyHRAcfxmBk10XUp8LhnimtcACnq4y3NcQ2CL90khDUTdwXKhlAc4bSXDZ79Zuo8rLSazzCwH61-iuyQgTLCaDtb5Ol4kJfPvswy2qOa1lbO1-RiIA8rMf_VkuRdBnKQ/s180-c/sori_snow_boy.png')"
      />
    </div>
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
</style>