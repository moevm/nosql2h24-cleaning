import { createWebHistory, createRouter } from 'vue-router'
import AuthenticationPage from '../components/AuthenticationPage.vue'
import MainPage from '../components/MainPage.vue'
import ClientMainPage from '../components/client/ClientMainPage.vue'
import ClientAddressesPage from '../components/client/ClientAddressesPage.vue'
import ClientOrdersHistoryPage from '../components/client/ClientOrdersHistoryPage.vue'
import ClientCreateOrderPage from '../components/client/ClientCreateOrderPage.vue'
import AdminMainPage from '../components/admin/AdminMainPage.vue'
import AdminWorkersSettingsPage from '../components/admin/AdminWorkersSettingsPage.vue'
import AdminServicesSettingsPage from '../components/admin/AdminServicesSettingsPage.vue'
import AdminStatisticSettingsPage from '../components/admin/AdminStatisticSettingsPage.vue'

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
        path: 'client:id([a-fA-F0-9]+)',
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
          },
          {
            name: 'client-create-order',
            path: 'create-order',
            component: ClientCreateOrderPage
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
          },
          {
            name: 'admin-statistic-settings',
            path: 'statistic',
            component: AdminStatisticSettingsPage
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
