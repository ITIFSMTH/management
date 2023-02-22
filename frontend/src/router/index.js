import VueCookies from 'vue-cookies';
import jwtDecode from "jwt-decode";
import { createRouter, createWebHistory } from "vue-router";
import { Roles } from '@/shared';
import Store from '@/store';


const routes = [
  {
    path: '/login',
    name: 'login-route',
    meta: {layout: 'login-layout'},
    component: () => import('@/views/Login.vue')
  },
  {
    path: '/admin',
    name: 'admin-route',
    meta: {layout: 'panel-layout', forRoles: [Roles.AdministratorRole]},
    component: () => import('@/views/Admin.vue')
  },
  {
    path: '/statistic',
    name: 'statistic-route',
    meta: {layout: 'panel-layout', forRoles: [Roles.AdministratorRole]},
    component: () => import('@/views/Statistic.vue')
  },
  {
    path: '/operator',
    name: 'operator-route',
    meta: {layout: 'panel-layout', forRoles: [Roles.JuniorOperatorRole, Roles.SeniorOperatorRole]},
    component: () => import('@/views/Operator.vue')
  },
  {
    path: '/operator/poll/rating',
    name: 'rating-poll-route',
    meta: {layout: 'panel-layout', forRoles: [Roles.JuniorOperatorRole]},
    component: () => import('@/views/RatingPoll.vue')
  },
  {
    path: '/operator/poll/budget',
    name: 'budget-poll-route',
    meta: {layout: 'panel-layout', forRoles: [Roles.SeniorOperatorRole]},
    component: () => import('@/views/BudgetPoll.vue')
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/admin'
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

router.beforeEach((to, from, next) => {
  // Get JWT cookie
  const jwtCookie = VueCookies.get('jwt')
  const jwtDecoded = jwtCookie ? jwtDecode(jwtCookie) : {}

  // If route don't need to be autrhorized
  if (!to.meta.forRoles || to.meta.forRoles.length === 0) return next()

  // Check jwt cookie valid
  if (!jwtCookie || Date.now() >= jwtDecoded.exp * 1000){
    return next({name: 'login-route'})
  }

  // Check is user role exist in route
  if (!to.meta.forRoles.find((role) => role.ID === jwtDecoded.role_id)) {
    if (Store.getters.isUserAdministrator) {
      return next({name: 'admin-route'})
    } else if (Store.getters.isUserOperator) {
      return next({name: 'operator-route'})
    } else {
      // Who the fuck you are?
      return
    }
  }

  // If no problems route user
  next()
})

export default router