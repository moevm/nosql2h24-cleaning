<script setup lang="ts">
import { defineProps } from 'vue';

const props = defineProps<{
  title?: string;
  height?: string;
  width?: string;
  items: any;
}>();
</script>

<template>
  <div
    class="header-container"
    :style="{
      height: props.height, 
      width: props.width 
    }"
  > <h1>{{ props.title }}</h1>
    <div class="items-container">
      <slot
        name="items"
        v-for="item in props.items"
        :key="item.id"
        :item="item"
      ></slot>
    </div>
    <slot
      name="bottom"
    ></slot>
  </div>
</template>

<style scoped>
.header-container {
  display: flex;
  flex-direction: column;
  text-align: center;
  height: 100%;
  width: 80%;
  background-color: white;
  border: 2px solid grey;
  border-radius: 25px;
}
.items-container {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-content: flex-start;
  flex-wrap: wrap;
  gap: 10px;
  padding-left: 10px;
  padding-right: 10px;
  overflow-y: scroll;
  scrollbar-width: none;
}
.items-container::-webkit-scrollbar {
  width: 0px;
  background: transparent;
}
h1 {
  font-size: 42px;
  font-weight: bold;
}
</style>
