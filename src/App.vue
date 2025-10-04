<script setup>
import { computed, ref } from 'vue'
import Generate from './components/Generate.vue'
import Home from './components/Home.vue'
import Train from './components/Train.vue'

const routes = {
  '/': Home,
  '/Train': Train,
  '/Generate': Generate
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
  <a href="#/Generate">キャラ生成(仮)</a> |
  <component :is="currentView" />
</template>

<style scoped>

</style>
