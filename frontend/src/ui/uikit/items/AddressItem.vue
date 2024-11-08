<script setup lang="ts">
import { defineProps } from 'vue'

const props = defineProps<{
  address: {
    id: string;
    address: string;
  }
}>()
</script>

<template>
  <div class="address-item">
    <p>{{ props.address.address }}</p>
  </div>
</template>

<style scoped>
.address-item {
  display: flex;
  flex-direction: row;
  justify-content: top;
  align-items: center;
  width: 49%;
  border: 3px solid #3846c0;
  border-radius: 15px;
  padding: 10px;
  text-align: left;
}
p {
  font-size: 32px;
  font-weight: bold;
}
</style>
