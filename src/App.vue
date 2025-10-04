<script setup>
import { computed, ref } from 'vue'
import Generate from './components/Generate.vue'
import Train from './components/Train.vue'
/* import NotFound from './NotFound.vue' */

const routes = {
  '/': Generate,
  '/Train': Train
}

const currentPath = ref(window.location.hash)

window.addEventListener('hashchange', () => {
  currentPath.value = window.location.hash
})

const currentView = computed(() => {
  return routes[currentPath.value.slice(1) || '/'] || NotFound
})
</script>

<template>
  <a href="#/">キャラ生成</a> |
  <a href="#/Train">キャラ育成</a> |
  <a href="#/non-existent-path">Broken Link</a>
  <component :is="currentView" />
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}

.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
