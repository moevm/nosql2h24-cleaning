<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Order from '../../../api/models/order'
import { getOrderById, getService } from '../../../api/request'
import Service from '../../../api/models/service'
import ActionButton from '../ActionButton.vue'

const route = useRoute()
const router = useRouter()
const orderId = route.params.order_id as string
const order = ref<Order | null>(null)
const services = ref<Service[]>([])

onMounted(() => {
  getOrderById(orderId)
    .then((response) => {
      order.value = response
      console.log(order.value)
      const date = new Date(order.value.date_time)
      order.value.date_time = date
      const servicePromises = order.value.services.map((serviceId: string) => getService(serviceId))
      return Promise.all(servicePromises)
    })
    .then((serviceResponses) => {
      services.value = serviceResponses;
    })
})

function formatDate(dateTime: string): string {
  const date = new Date(dateTime)
  const month = date.getMonth() + 1 < 10 ? `0${date.getMonth() + 1}` : date.getMonth() + 1
  const minutes = date.getMinutes() < 10 ? `0${date.getMinutes()}` : date.getMinutes()
  return `${date.getDate()}.${month}.${date.getFullYear()} ${date.getHours()}:${minutes}`
}

function goBack() {
  router.push({ name: 'client-orders-history' });
}

</script>

<template>
  <div v-if="order" class="order-detail">
    <h1>Детали заказа</h1>
    <p>Заказ по адресу: {{ order!.address.city }}, {{ order!.address.street }}, дом {{ order!.address.building }}, подъезд {{ order!.address.entrance }}, этаж {{ order!.address.floor }}, квартира {{ order!.address.door_number }}</p>
    
    <p>Информация о помещении: </p>
    <ul>
      <li>Площадь: {{ order!.area }} м²</li>
      <li>Количество комнат: {{ order!.number_of_rooms }}</li>
      <li>Количество санузлов: {{ order!.number_of_bathrooms }}</li>
      <li>Загрязненность: {{ order!.pollution }}</li>
    </ul>
    <p>На дату: {{ formatDate(order!.date_time.toISOString()) }}</p>
    <p>Статус заказа: {{ order!.status }}</p>
    <p>Цена: {{ order!.price }} ₽</p>
    <p>Количество работников: {{ order!.required_workers }}</p>
    <p>Услуги:</p>
    <p v-for="service in services" :key="service.id">- {{ service.name }}</p>
    <ActionButton
      text="Назад"
      color="red"
      :onClick="goBack"
    ></ActionButton>
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
  font-weight: bold;
}
p {
  font-size: 28px;
}
ul {
	margin: 0 0 0 30px;
}
ul li {
  font-size: 24px;
}
</style>
