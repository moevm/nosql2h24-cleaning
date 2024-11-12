<script setup lang="ts">
import { inject, onMounted } from 'vue'
import { RouterView } from 'vue-router'
import { useUserStore } from '../../store/user';
import { getUsers } from '../../api/request';

const userStore = useUserStore()
const user = userStore.getUser()
const addSideBarButtons: Function | undefined = inject('addSideBarButtons')
const setUserCard: Function | undefined = inject('setUserCard')

onMounted(() => {
  addSideBarButtons!(
    {
      text: 'Заказы',
      type: 'orders',
      to: `/cleaning/admin/orders`
    }
  )
  addSideBarButtons!(
    {
      text: 'Исполнители',
      type: 'workers',
      to: `/cleaning/admin/workers`
    }
  )
  addSideBarButtons!(
    {
      text: 'Услуги',
      type: 'services',
      to: `/cleaning/admin/services`
    }
  )
  addSideBarButtons!(
    {
      text: 'Статистика',
      type: 'statistics',
      to: `/cleaning/admin/statistics`
    }
  )
  setUserCard!(`${user?.name} ${user?.surname}`, `${user?.email}`)
  
  getUsers('CLIENT').then((users) => {
    console.log(users)
  })
})
</script>

<template>
  <RouterView />
</template>

<style scoped>
</style>
