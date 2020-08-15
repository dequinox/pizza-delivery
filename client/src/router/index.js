import Vue from 'vue'
import Router from 'vue-router'
import Register from '@/components/Register'
import Login from '@/components/Login'
import Pizzas from '@/components/Pizzas'
import Order from '@/components/Order'
import History from '@/components/History'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/register',
      name: 'register',
      component: Register
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/order',
      name: 'order',
      component: Order
    },
    {
      path: '/pizzas',
      name: 'pizzas',
      component: Pizzas
    },
    {
      path: '/history',
      name: 'history',
      component: History
    },
    {
      path: '*',
      redirect: 'pizzas'
    },
  ]
})
