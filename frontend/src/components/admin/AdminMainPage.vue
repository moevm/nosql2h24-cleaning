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
      to: `/cleaning/admin/statistics`
    }
  )

  getUserInfo('me').then((user) => {
    // userStore.setUser(user)
    setUserCard!(`${user.name} ${user.surname}`, `${user.email}`)
  })
  
})

getUsers('CLIENT').then((response) => {
  clients.value = response
})
getUsers('WORKER').then((response) => {
  workers.value = response
})
</script>

<template>
  <RouterView />
  <div>
    <v-data-table
      :headers="headers"
      :items="clients"
      class="elevation-1"
    >
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>Клиенты</v-toolbar-title>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>
        </v-toolbar>
      </template>
      <template v-slot:item.name="{ item }">
        <span>{{ item.name }}</span>
      </template>
      <template v-slot:item.email="{ item }">
        <span>{{ item.email }}</span>
      </template>
      <!-- Добавьте другие слоты для отображения данных по мере необходимости -->
    </v-data-table>
  </div>
  <div>
    <v-data-table
      :headers="headers"
      :items="workers"
      class="elevation-1"
    >
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>Исполнители</v-toolbar-title>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>
        </v-toolbar>
      </template>
      <template v-slot:item.name="{ item }">
        <span>{{ item.name }}</span>
      </template>
      <template v-slot:item.email="{ item }">
        <span>{{ item.email }}</span>
      </template>
      <!-- Добавьте другие слоты для отображения данных по мере необходимости -->
    </v-data-table>
  </div>
</template>

<style scoped>
</style>
