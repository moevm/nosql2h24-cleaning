<script setup lang="ts">
import { inject, ref, onMounted } from 'vue'
import {
  RouteLocationNormalizedLoadedGeneric,
  useRoute, 
  RouterView 
} from 'vue-router'
import { getUserInfo } from '../../api/request'

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
  getUserInfo('me').then((user) => {
    setUserCard!(`${user.name} ${user.surname}`, `${user.email}`)
  })
  .catch((error) => {
    console.error("Failed to fetch user info:", error)
  })

})
</script>

<template>
  <RouterView />
</template>

<style scoped>
</style>
