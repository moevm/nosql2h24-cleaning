<script setup lang="ts">
import { defineProps, withDefaults } from 'vue'

const props = withDefaults(defineProps<{
  flexDirection?: string;
  justifyContent?: string;
  alignItems?: string;
}>(), {
  flexDirection: 'column',
  justifyContent: 'center',
  alignItems: 'center'
})

</script>

<template>
  <div
    class="main-container"
    :style="{
      flexDirection: props.flexDirection,
      justifyContent: props.justifyContent,
      alignItems: props.alignItems
    }"
  >
    <slot></slot>
  </div>
</template>

<style scoped>
.main-container {
  max-width: 100%;
  max-height: 100%;
  display: flex;
  height: 100%;
  width: 80%;
  background-color: white;
  gap: 20px;
  padding-top: 20px;
  padding-bottom: 20px;
}
</style>
