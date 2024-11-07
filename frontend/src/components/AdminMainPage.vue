<script setup lang="ts">
import { inject, onMounted } from 'vue'
import { RouterView } from 'vue-router'

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
  setUserCard!('Иван Иванов', 'example@mail.com')
})
</script>

<template>
  <RouterView />
</template>

<style scoped>
</style>
