<script setup lang="ts">
import { defineProps } from 'vue';

const props = defineProps<{
  service: {
    name: string;
    price: number;
  };
}>();

function addToCart(): void {
  console.log(`Added to cart`);
}
</script>

<template>
  <div class="service-item">
    <div class="text">
      <div class="name">{{ props.service.name }}</div>
      <div class="price">{{ props.service.price }}$</div>
    </div>
    <button class="add-button" @click="addToCart">
      <span>+</span>
    </button>
  </div>
</template>

<style scoped>
.service-item {
  display: flex;
  align-items: start;
  justify-content: space-between;
  border: 1px solid #ddd;
  border-radius: 15px;
  padding: 10px;
  width: 500px;
  font-family: Arial, sans-serif;
}

.text {
  display: flex;
  flex-direction: column;
}

.title {
  font-size: 16px;
}

.price {
  font-size: 14px;
  color: black;
  align-items: start;
}

.add-button {
  background-color: #394cc2;
  color: white;
  border: none;
  border-radius: 50%;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  cursor: pointer;
}

</style>