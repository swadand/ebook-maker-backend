<script setup lang="ts">

</script>

<template>
  <div class="menu-container">
    <div class="logo">EM</div>
    <div>
      <div class="file-name">
        Name of File
      </div>
      <div class="menu">
        <div class="menu-items">File</div>
        <div class="menu-items">View</div>
        <div class="menu-items">Insert</div>
        <div class="menu-items">About</div>
        <div class="menu-items">Help</div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.logo {
  margin: 8px;
  padding: 10px;
  position: relative;
  width: 25px;
  background-color: rgb(206, 206, 206);
  border-radius: 100%;
}

.file-name {
  display: flex;
  margin-left: 10px;
  margin-top: 5px;
}

.menu-container {
  z-index: 50;
  position: sticky;
  background-color: rgb(236, 236, 236);
  top: 0;
  display: flex;
  gap: 2px;
}
.menu {
  display: flex;
  gap: 5px;
}
.menu-items {
  font-size: small;
  background-color: rgb(238, 238, 238);
  border-radius: 4px;
  width: 28%;
  padding: 5px;
  user-select: none;
}
.menu-items:hover {
  background-color: rgb(221, 221, 221);
  cursor: pointer;
}
</style>
