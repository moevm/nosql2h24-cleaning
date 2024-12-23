<script setup lang="ts">
import { defineProps, ref } from 'vue';
import Order from '../../api/models/order'
import ActionButton from '../../ui/uikit/ActionButton.vue'

const currentDate = ref<Date>(new Date());

setInterval(() => {
  currentDate.value = new Date();
}, 1000 * 60);


const props = defineProps<{
  order: Order;
  type: string;
  callback: (id: string) => void;
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

async function btnClick() {
  props.callback(props.order.id ?? '');
}

</script>

<template>
  <div class="order-item">
    <p>{{ formatDate() }}</p>
    <p>{{ formatAddress() }}, </p>
    <div class="order-price">
      <p>{{ props.order.price }} ₽</p>
      <ActionButton
        v-if="props.type === 'TAKE'" 
        text="Взять заказ"
        type="add"
        color="#394cc2"
      @click="btnClick">
    </ActionButton>
    <ActionButton
        v-if="props.type === 'CONFIRM' && (props.order.status === 'CREATED' || props.order.status === 'IN_PROGRESS') && currentDate >= new Date(props.order.date_time)" 
        text="Завершить заказ"
        type="add"
        color="#394cc2"
      @click="btnClick">
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
