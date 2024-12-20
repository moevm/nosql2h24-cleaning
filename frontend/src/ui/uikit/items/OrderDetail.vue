<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import Order from '../../../api/models/order'
import { getOrderById } from '../../../api/request';

const route = useRoute();
const orderId = route.params.id as string;
const order = ref<Order>();

onMounted(() => {
  getOrderById(orderId)
    .then((response) => {
      order.value = response
      console.log(order.value)
    })
})

function formatDate(dateTime: string): string {
  const date = new Date(dateTime);
  const month = date.getMonth() + 1 < 10 ? `0${date.getMonth() + 1}` : date.getMonth() + 1;
  const minutes = date.getMinutes() < 10 ? `0${date.getMinutes()}` : date.getMinutes();
  return `${date.getDate()}.${month}.${date.getFullYear()} ${date.getHours()}:${minutes}`;
}

</script>

<template>
  <div class="order-detail">
    <h1>Детали заказа</h1>
    <p>Заказ по адресу: {{ order!.address.city }}, {{ order!.address.street }}, дом {{ order!.address.building }}, подъезд {{ order!.address.entrance }}, этаж {{ order!.address.floor }}, квартира {{ order!.address.door_number }}</p>
    <p>На дату: {{ formatDate(order!.date_time.toISOString()) }}</p>
    <p>Цена: {{ order!.price }}₽</p>
    <p>Количество работников: {{ order!.required_workers }}</p>
    <p>Услуги:</p>
  </div>
</template>

<style scoped>
.order-detail {
  border: 3px solid #3846c0;
  width: 100%;
  height: 100%;
  border-radius: 15px;
  padding: 10px;
  text-align: left;
}
</style>
