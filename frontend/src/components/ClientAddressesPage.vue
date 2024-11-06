<script setup lang="ts">
import { ref } from 'vue'
import HeaderList from '../ui/uikit/HeaderList.vue'
import AddressItem from '../ui/uikit/AddressItem.vue'

const address = ref([
  {id: 1, address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4"},
  {id: 2, address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4"},
  {id: 3, address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4"},
  {id: 4, address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4"},
  {id: 5, address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4"}
]) // TODO DB request
</script>

<template>
  <div class="container">
    <HeaderList
      title="Мои адреса" 
      :items="address"
      height="100%"
      width="95%"
    >
      <template #items="{ item }">
        <AddressItem :address="item" />
      </template>
    </HeaderList>
  </div>
</template>

<style scoped>
.container {
  max-width: 100%;
  max-height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 80%;
  background-color: white;
  gap: 20px;
  padding-top: 20px;
  padding-bottom: 20px;
}
</style>
