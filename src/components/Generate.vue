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