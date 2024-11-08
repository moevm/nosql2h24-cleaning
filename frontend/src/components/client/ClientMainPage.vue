<script setup lang="ts">
import { inject, onMounted } from 'vue'
import {
  RouteLocationNormalizedLoadedGeneric,
  useRoute, 
  RouterView 
} from 'vue-router'

const route: RouteLocationNormalizedLoadedGeneric = useRoute()

const addSideBarButtons: Function | undefined = inject('addSideBarButtons')
const setUserCard: Function | undefined = inject('setUserCard')

const clientId: string | string[] = route.params.id

onMounted(() => {
  addSideBarButtons!(
    {
      text: 'Создать новый заказ',
      type: 'create-order',
      to: `/cleaning/client${clientId}/create-order`
    }
  )
  addSideBarButtons!(
    {
      text: 'История заказов',
      type: 'history-order',
      to: `/cleaning/client${clientId}/history-order`
    }
  )
  addSideBarButtons!(
    {
      text: 'Мои адреса',
      type: 'my-addresses',
      to: `/cleaning/client${clientId}/my-addresses`
    }
  )
  setUserCard!('Иван Иванов', 'example@mail.com')
})
</script>

<template>
  <RouterView />
</template>

<style scoped>
</style>
