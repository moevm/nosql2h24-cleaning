import { createWebHistory, createRouter } from 'vue-router'
import AuthenticationPage from '../components/AuthenticationPage.vue'
import MainPage from '../components/MainPage.vue'
import ClientMainPage from '../components/ClientMainPage.vue'
import OrdersHistoryPage from '../components/OrdersHistoryPage.vue'

const routes = [
  {
    name: 'login',
    path: '/login',
    component: AuthenticationPage
  },
  {
    name: 'cleaning',
    path: '/cleaning',
    component: MainPage,
    children: [
      {
        name: 'client',
        path: 'client:id(\\d+)',
        component: ClientMainPage
      }
    ]
  },
  {
    name: 'orders',
    path: '/orders',
    component: OrdersHistoryPage
  },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})
  
export default router;
