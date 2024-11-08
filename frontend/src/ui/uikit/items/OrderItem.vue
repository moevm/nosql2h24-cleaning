<script setup lang="ts">
import { defineProps } from 'vue';

const props = defineProps<{
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
    <p>Заказ {{ props.order.date }} в {{ props.order.time }}</p>
    <p>{{ props.order.address }}</p>
    <p>{{ props.order.price }}$</p>
  </div>
</template>

<style scoped>
.order-item {
  border: 3px solid #3846c0;
  width: 100%;
  border-radius: 15px;
  padding: 10px;
  text-align: left;
}
</style>
