<script setup lang="ts">
import { ref, onMounted } from 'vue'
import PanelContainer from '../../ui/uikit/containers/PanelContainer.vue'
import HeaderList from '../../ui/uikit/containers/HeaderList.vue'
import OrderItem from '../../ui/uikit/items/OrderItem.vue'
import ActionButton from '../../ui/uikit/ActionButton.vue'
import DateFilter from '../../ui/uikit/DateFilter.vue'
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import { getAllOrders, getUserInfo } from '../../api/request'
import Order from '../../api/models/order'

const orders = ref<Order[]>([])

function fetchAllOrders() {
  getAllOrders('me')
    .then((response) => {
      orders.value = response
    })
    .catch((error) => {
      console.error('Fetch all orders error:', error);
    });
}

onMounted(() => {
  fetchAllOrders()
})

</script>

<template>
  <MainContainer>
    <PanelContainer>
      <ActionButton
        text="Все"
        color="#394cc2"
      ></ActionButton>
      <DateFilter />
      <ActionButton
        text="Поиск"
        color="#394cc2"
      ></ActionButton>
    </PanelContainer>
    <HeaderList 
      title="История заказов" 
      :items="orders"
      height="92%"
      width="95%"
      >
      <template #items="{ item }">
        <OrderItem :order="item" />
      </template>
    </HeaderList>
  </MainContainer>
</template>

<style scoped>
</style>
