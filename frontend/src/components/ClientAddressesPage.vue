<script setup lang="ts">
import { ref } from 'vue'
import HeaderList from '../ui/uikit/HeaderList.vue'
import AddressItem from '../ui/uikit/AddressItem.vue'
import MainContainer from '../ui/uikit/MainContainer.vue'

const address = ref([
  {id: 1, address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4"},
  {id: 2, address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4"},
  {id: 3, address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4"},
  {id: 4, address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4"},
  {id: 5, address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4"}
]) // TODO DB request
</script>

<template>
  <MainContainer>
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
  </MainContainer>
</template>

<style scoped>
</style>
