<script setup lang="ts">
import { defineProps } from 'vue'
import Order from '../../../api/models/order'
import { useRouter } from 'vue-router'

const router = useRouter()

const props = defineProps<{
  order: Order;
}>()

function formatDate(): string {
  const date = new Date(props.order.date_time)
  const mounth = date.getMonth() + 1 < 10 ? `0${date.getMonth() + 1}` : date.getMonth() + 1
  const minutes = date.getMinutes() < 10 ? `0${date.getMinutes()}` : date.getMinutes()
  return `${date.getDate()}.${mounth}.${date.getFullYear()} ${date.getHours()}:${minutes}`
}

function goToOrderDetails() {
  console.log('goToOrderDetails')
  router.push({ name: 'order-details', params: { order_id: props.order.id } });
}
</script>

<template>
  <div 
    class="order-item"
    @click="goToOrderDetails"
  >
    <p>
        Заказ по адресу: {{ props.order.address.city }}, 
        {{ props.order.address.street }}, 
        дом {{ props.order.address.building }},
        подъезд {{ props.order.address.entrance }},
        этаж {{ props.order.address.floor }},
        квартира {{ props.order.address.door_number }}
      </p>
      <p>
        На дату: {{ formatDate() }}</p>
      <p>
        Цена: {{ props.order.price }}₽
      </p>
    </div>
</template>

<style scoped>
.order-item {
  border: 3px solid #3846c0;
  width: 100%;
  border-radius: 15px;
  padding: 10px;
  text-align: left;
  cursor: pointer;
  transition: background-color 0.3s;
}
.order-item:hover {
  background-color: #f0f0f0;
}
p {
  font-size: 22px;
  font-weight: bold;
}
</style>
