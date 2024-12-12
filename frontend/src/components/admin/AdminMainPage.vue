<script setup lang="ts">
import { inject, onMounted } from 'vue'
import { RouterView } from 'vue-router'
import { getUserInfo, getUsers } from '../../api/request'
import { ref } from 'vue' 
import { User } from '../../api/models/user'

// const userStore = useUserStore()
// const user = userStore.getUser()

const clients = ref<User[]>([]);
const workers = ref<User[]>([]);
const headers = [
  { text: 'ID', value: 'id' },
  { text: 'Имя', value: 'name' },
  { text: 'Email', value: 'email' },
  // Добавьте другие заголовки по мере необходимости
];

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
