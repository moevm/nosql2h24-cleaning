import { createWebHistory, createRouter } from 'vue-router'
import AuthenticationPage from '../components/AuthenticationPage.vue';

const routes = [
  {
    name: "login",
    path: "/login",
    component: AuthenticationPage
  }
];

const router = createRouter({
    history: createWebHistory(),
    routes,
})
  
export default router;
