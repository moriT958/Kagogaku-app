import vue from "@vitejs/plugin-vue";
import path from 'path';
import { defineConfig } from "vite";

// https://vite.dev/config/
export default defineConfig({
  base: '/Kagogaku-app/',
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve('resources/js'),
      $image: path.resolve('public/img'),
    },
  },
  server: {
    port: 3000,
    host: "0.0.0.0",
    proxy: {
      "/api": {
        target: "http://localhost:8080",
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ""),
      },
    },
  },
});
