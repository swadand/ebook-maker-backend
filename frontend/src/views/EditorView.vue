<script setup lang="ts">
  import { ref } from "vue";
  import Editor from "../components/Editor.vue";
  import MenuBar from "../components/MenuBar.vue";
  
  const content = ref("");
</script>

<template>
  <MenuBar />
  <Editor />
</template>
