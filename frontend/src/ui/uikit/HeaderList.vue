<script setup lang="ts">
import OrderItem from "./OrderItem.vue";
import { defineProps } from 'vue';


const props = defineProps<{
  title: string;
  orders: Array<{
    id: number;
    date: string;
    time: string;
    address: string;
    price: number;
  }>;
}>();
</script>

<template>
  <div class="header-list">
    <h1>{{ title }}</h1>
    <div class="orders-list">
      <OrderItem v-for="order in orders" :key="order.id" :order="order" />
    </div>
  </div>
</template>

<style scoped>
.header-list {
  display: flex;
  flex-direction: column;
  height: 90vh;
  box-sizing: border-box;
  padding-left: 10px;
  padding-right: 10px;
  border: 2px solid grey;
  border-radius: 25px;
  text-align: center;
}

.title {
  flex-shrink: 0;
  font-size: 24px;
  margin-bottom: 10px;
}

.orders-list {
  display: grid;
  gap: 10px;
  flex-grow: 1;
  overflow-y: scroll;
  scrollbar-width: none;
}

.orders-list::-webkit-scrollbar {
  width: 0px;
  background: transparent;
}
</style>