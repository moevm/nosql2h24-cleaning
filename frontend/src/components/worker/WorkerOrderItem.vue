<script setup lang="ts">
import { defineProps } from 'vue';
import Order from '../../api/models/order'
import ActionButton from '../../ui/uikit/ActionButton.vue';
import { updateOrder } from '../../api/request'

const emit = defineEmits(['update-orders']);

const props = defineProps<{
  order: Order;
  workerId?: string
}>();

function formatAddress(): string {
  return `г. ${props.order.address.city}, ул. ${props.order.address.street}, д. ${props.order.address.building}, 
  кв. ${props.order.address.door_number}, п. ${props.order.address.entrance}, э. ${props.order.address.floor}`
}

function formatDate(): string {
  const date = new Date(props.order.date_time)
  const mounth = date.getMonth() + 1 < 10 ? `0${date.getMonth() + 1}` : date.getMonth() + 1
  const minutes = date.getMinutes() < 10 ? `0${date.getMinutes()}` : date.getMinutes()
  return `${date.getDate()}.${mounth}.${date.getFullYear()} ${date.getHours()}:${minutes}`
}

async  function takeOrder() {
    const updatedOrder = { // тут пока что непонятно какую логику юзаем и на фронте ли это должно лежать
        ...props.order,
        status: 'INPROGRESS',
        workers: [...(props.order.workers ?? []), props.workerId]
    };
    await updateOrder(updatedOrder as Order)
    .then(_ => {
        emit('update-orders')
        console.log('Worker get order');
    }).catch((error) => {
        console.log('Worker get order error:', error);
    })
}

</script>

<template>
  <div class="order-item">
    <p>{{ formatDate() }}</p>
    <p>{{ formatAddress() }}, </p>
    <div class="order-price">
      <p>{{ props.order.price }}$</p>
      <ActionButton 
        text="+"
        type="add"
        color="#394cc2"
      @click="takeOrder">
    </ActionButton>
    </div>
  </div>
</template>

<style scoped>
.order-item {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  width: 49%;
  border: 3px solid #3846c0;
  border-radius: 15px;
  padding: 10px;
  text-align: left;
}
.order-price {
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
.order-status {
  width: 30%;
  display: flex;
  justify-content: center;
  border-radius: 15px;
}
.order-status-in-progress {
  border: 3px solid #f7b500;
}
.order-status-created {
  border: 3px solid #f70000;
}
.order-status-finished {
  border: 3px solid #04f700;
}
p {
  font-size: 22px;
  font-weight: bold;
}
</style>
