import { createWebHistory, createRouter } from 'vue-router'
import AuthenticationPage from '../components/AuthenticationPage.vue'
import MainPage from '../components/MainPage.vue'
import ClientMainPage from '../components/ClientMainPage.vue'
import ClientAddressesPage from '../components/ClientAddressesPage.vue'
import ClientOrdersHistoryPage from '../components/ClientOrdersHistoryPage.vue'
import AdminMainPage from '../components/AdminMainPage.vue'
import AdminWorkersSettingsPage from '../components/AdminWorkersSettingsPage.vue'
import AdminServicesSettingsPage from '../components/AdminServicesSettingsPage.vue'

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
        component: ClientMainPage,
        children: [
          {
            name: 'client-addresses',
            path: 'my-addresses',
            component: ClientAddressesPage
          },
          {
            name: 'client-orders-history',
            path: 'history-order',
            component: ClientOrdersHistoryPage
          }
        ]
      },
      {
        name: 'admin',
        path: 'admin',
        component: AdminMainPage,
        children: [
          {
            name: 'admin-workers-settings',
            path: 'workers',
            component: AdminWorkersSettingsPage
          },
          {
            name: 'admin-services-settings',
            path: 'services',
            component: AdminServicesSettingsPage
          }
        ]
      }
    ]
  }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})
  
export default router;
