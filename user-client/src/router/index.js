import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: () => import('@/views/Login.vue')
    },
    {
      path: '/',
      name: 'Root',
      component: () => import('@/views/Root.vue'),
      redirect: '/home',
      children: [
        {
          path: '/home',
          name: 'Home',
          component: () => import('@/views/Home.vue')
        },
        {
          path: '/south',
          name: 'South',
          component: () => import('@/views/SouthFruit.vue')
        },
        {
          path: '/north',
          name: 'North',
          component: () => import('@/views/NorthFruit.vue')
        },
        {
          path: 'fruits',
          name: 'Fruits',
          component: () => import('@/views/Fruits.vue')
        },
        {
          path: '/order',
          name: 'Order',
          component: () => import('@/views/Order.vue')
        },
        {
          path: '/cart',
          name: 'Cart',
          component: () => import('@/views/Cart.vue')
        }
      ]
    },
    {
      path: '/pay',
      name: 'Pay',
      component: () => import('@/views/Pay.vue')
    }
  ]
})

export default router
