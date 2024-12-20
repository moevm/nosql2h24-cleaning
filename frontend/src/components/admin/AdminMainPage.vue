<script setup lang="ts">
import { inject, onMounted } from 'vue'
import { RouterView } from 'vue-router'
import { getUserInfo } from '../../api/request'

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
      to: `/cleaning/admin/statistic`
    }
  )

  getUserInfo('me').then((user) => {
    setUserCard!(`${user.name} ${user.surname}`, `${user.email}`)
  })
  
})
</script>

<template>
  <RouterView />
</template>

<style scoped>
</style>
