<script setup lang="ts">
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import PanelContainer from '../../ui/uikit/containers/PanelContainer.vue'
import HeaderList from '../../ui/uikit/containers/HeaderList.vue'
import WorkerOrderItem from '../worker/WorkerOrderItem.vue'
import { onMounted, ref } from 'vue'
import { getFiltredOrders, getUserInfo, takeOrder } from '../../api/request';
import Order from '../../api/models/order'
import { useUserStore } from '../../store/user'
const orders = ref<Order[]>([]);
const store = useUserStore();

function fetchAllAvailableOrders() : Promise<void>  {
  return getFiltredOrders({
    statuses: [ 'CREATED' ]
  }).then((response) => {
    response = response.filter((order) => {
      return !order.workers?.includes(store.getUser()?.id ?? '');
    });
    orders.value = response
  }).catch((error) => {
    console.error('Couldnt load orders:', error);
  });
}

function fetchWorkerId() : Promise<void> {
  return getUserInfo('me')
  .then((response) => {
    store.setUser(response);
  })
  .catch((error) => {
    console.error("Error get worker info:", error)
  })
}

function takeOrderCallback(order_id: string) {
  takeOrder(order_id).then(() => {
    fetchAllAvailableOrders();
  });
}

onMounted(async () => {
  await fetchWorkerId();
  await fetchAllAvailableOrders();
})
</script>

<template>
  <MainContainer>
    <HeaderList
      title="Доступные заказы" 
      :items="orders"
      height="100%"
      width="90%"
    >
      <template #items="{ item }">
        <WorkerOrderItem :order="item" :callback="takeOrderCallback" :type="'TAKE'"/>
      </template>
      <template #empty>
        <h1>Нет доступных заказов</h1>
      </template>
    </HeaderList>
  </MainContainer>
</template>