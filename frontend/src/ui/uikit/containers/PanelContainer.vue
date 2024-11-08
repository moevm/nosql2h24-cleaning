<script setup lang="ts">
import { defineProps, withDefaults } from 'vue'

const props = withDefaults(defineProps<{
  height?: string,
  width?: string
}>(), {
  height: '8%',
  width: '95%'
})
</script>

<template>
  <div
    class="panel-container"
    :style="{
      height: props.height,
      width: props.width
    }"
  >
    <slot></slot>
  </div>
</template>

<style scoped>
.panel-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  /* height: 8%;
  width: 95%; */
  gap: 10px;
}
</style>
