<script setup lang="ts">
</script>

<template>
  <div class="main-container">
    <slot></slot>
  </div>
</template>

<style scoped>
.main-container {
  max-width: 100%;
  max-height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 80%;
  background-color: white;
  gap: 20px;
  padding-top: 20px;
  padding-bottom: 20px;
}
</style>
