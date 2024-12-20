<script setup lang="ts">
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import PanelContainer from '../../ui/uikit/containers/PanelContainer.vue'
import ActionButton from '../../ui/uikit/ActionButton.vue'
import HeaderList from '../../ui/uikit/containers/HeaderList.vue'
import WorkerOrderItem from '../worker/WorkerOrderItem.vue'
import { onMounted, ref } from 'vue'
import { getAllOrders, getUserInfo } from '../../api/request';
import Order from '../../api/models/order'

const orders = ref<Order[]>([]);
const workerId = ref<string>()

async function fetchAllOrders() {
  getAllOrders(null)
  .then((response) => {
    const filteredOrders = response.filter((order) => {
      const workersCount = order.workers?.length ?? 0;
      return workersCount < order.required_workers;
    });
    orders.value = filteredOrders
  })
  .catch((error) => {
    console.error('Couldnt load orders:', error);
  });
}

async function fetchWorkerId() {
  getUserInfo('me')
  .then((response) => {
    workerId.value = response.id
  })
  .catch((error) => {
    console.error("Error get worker info:", error)
  })
}

onMounted(() => {
  fetchWorkerId()
  fetchAllOrders()
})
</script>

<template>
  <MainContainer>
    <PanelContainer
      height="10%"
      width="95%"
    >
      <div class="filters-btn">
        <ActionButton
          text="Все"
          type="all"
          color="#394cc2"
        ></ActionButton>
      </div>
    </PanelContainer>
      <HeaderList
        title="Заказы" 
        :items="orders"
        height="100%"
        width="95%"
      >
        <template #items="{ item }">
          <WorkerOrderItem :order="item" :workerId="workerId" @update-orders="fetchAllOrders"/>
        </template>
      </HeaderList>
  </MainContainer>
</template>