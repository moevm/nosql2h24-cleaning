<script setup lang="ts">
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import PanelContainer from '../../ui/uikit/containers/PanelContainer.vue'
import HeaderList from '../../ui/uikit/containers/HeaderList.vue'
import WorkerOrderItem from '../worker/WorkerOrderItem.vue'
import { onMounted, ref } from 'vue'
import { completeOrder, getFiltredOrders, getUserInfo } from '../../api/request';
import Order from '../../api/models/order'
import { useUserStore } from '../../store/user'

const orders = ref<Order[]>([]);
const store = useUserStore();

function fetchAllMyOrders(): Promise<void> {
  return getFiltredOrders({
    statuses: [ 'CREATED', 'IN_PROGRESS' ],
    workers_id: [ store.getUser()?.id ?? '' ],
  }).then((response) => {
    response = response.filter((order) => {
      return order.workers?.includes(store.getUser()?.id ?? '');
    });
    orders.value = response
  }).catch((error) => {
    console.error('Couldnt load orders:', error);
  });
}

function fetchWorkerId(): Promise<void> {
  return getUserInfo('me')
  .then((response) => {
    store.setUser(response);
  })
  .catch((error) => {
    console.error("Error get worker info:", error)
  })
}

function completeOrderCallback(order_id: string) {
  completeOrder(order_id).then(() => {
    fetchAllMyOrders();
  });
}

onMounted(async () => {
    await fetchWorkerId();
    await fetchAllMyOrders();
})
</script>

<template>
  <MainContainer>
    <HeaderList
      title="Взятые заказы" 
      :items="orders"
      height="100%"
      width="90%"
    >
      <template #items="{ item }">
        <WorkerOrderItem :order="item" :callback="completeOrderCallback" :type="'CONFIRM'"/>
      </template>
      <template #empty>
        <h1>У вас нет взятых заказов</h1>
      </template>
    </HeaderList>
  </MainContainer>
</template>