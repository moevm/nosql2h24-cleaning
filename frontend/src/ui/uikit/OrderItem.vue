<script setup lang="ts">
import { defineProps } from 'vue';

const { order } = defineProps<{
  order: {
    date: string;
    time: string;
    address: string;
    price: number;
  };
}>();
</script>

<template>
  <div class="order-item">
    <p>Заказ {{ order.date }} в {{ order.time }}</p>
    <p>{{ order.address }}</p>
    <p>{{ order.price }}$</p>
  </div>
</template>

<style scoped>
.order-item {
  border: 3px solid #3846c0;
  border-radius: 15px;
  padding: 10px;
  text-align: left;
}
</style>
