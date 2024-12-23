<script setup lang="ts">
import { inject, onMounted } from 'vue'
import {
  RouteLocationNormalizedLoadedGeneric,
  useRoute, 
  RouterView 
} from 'vue-router'
import { getUserInfo } from '../../api/request'
import { useUserStore } from '../../store/user';

const route: RouteLocationNormalizedLoadedGeneric = useRoute()
const addSideBarButtons: Function | undefined = inject('addSideBarButtons')
const setUserCard: Function | undefined = inject('setUserCard')

const clientId: string | string[] = route.params.id

onMounted(() => {
  addSideBarButtons!(
    {
      text: 'Cписок доступных заказов',
      type: 'orders',
      to: `/cleaning/worker${clientId}/orders`
    }
  )
  addSideBarButtons!(
    {
      text: 'Мои заказы',
      type: 'taked-orders',
      to: `/cleaning/worker${clientId}/taked-orders`
    }
  )

  getUserInfo('me').then((user) => {
    useUserStore().setUser(user);
    setUserCard!(`${user.name} ${user.surname}`, `${user.email}`);
  })
})
</script>

<template>
  <RouterView />
</template>