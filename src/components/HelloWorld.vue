<script setup>
import { ref, onMounted } from "vue";

// 親コンポーネント(SFC)から受け取るデータ
// 今は App.vue が親にあたる
defineProps({
  msg: String,
});

const helloMsg = ref(""); // バックエンドからのデータ
const loading = ref(false); // ローディング中かどうか
const error = ref(""); // エラー

// バックエンド API を叩くサンプル
const fetchHello = async () => {
  loading.value = true;
  error.value = "";
  try {
    // GET /hello エンドポイントを叩く
    const response = await fetch("/api/hello");
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();
    helloMsg.value = data.message;
  } catch (err) {
    error.value = `エラー: ${err.message}`;
    console.error("Failed to fetch from backend:", err);
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <h1>{{ msg }}</h1>

  <div class="card">
    <h2>Hello, World!!🌎</h2>
    <p v-if="loading">読み込み中...</p>
    <p v-else-if="error" style="color: red">{{ error }}</p>
    <p v-else-if="helloMsg" style="color: #42b883; font-weight: bold">
      {{ helloMsg }}
    </p>
    <button type="button" @click="fetchHello">バックエンド API を叩く</button>
  </div>
</template>

<style scoped></style>
