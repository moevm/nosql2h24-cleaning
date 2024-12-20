<script setup lang="ts">
import { ref } from 'vue'
import PanelContainer from '../../ui/uikit/containers/PanelContainer.vue'
import HeaderList from '../../ui/uikit/containers/HeaderList.vue'
import OrderItem from '../../ui/uikit/items/OrderItem.vue'
import ActionButton from '../../ui/uikit/ActionButton.vue'
import DateFilter from '../../ui/uikit/DateFilter.vue'
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'

const orders = ref([
  { id: 1, date: "28.02.2024", time: "10:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 37 },
  { id: 2, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
  { id: 4, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
  { id: 5, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
  { id: 6, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
  { id: 7, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
  { id: 8, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
  { id: 9, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
  { id: 10, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
  { id: 11, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
  { id: 12, date: "15.12.2023", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 42 }
]) // TODO DB request
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
