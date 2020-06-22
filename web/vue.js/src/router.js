import jwtDecode from 'jwt-decode';
import Vue from 'vue';
import Router from 'vue-router';

import { AUTH_TOKEN } from './constants';
import { IJWTDecoded } from './store/modules/user';

Vue.use(Router);

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {path: '/', name: 'home', component: () => import(/* webpackChunkName: "home" */ './views/Home.vue')},
    {path: '/privacy', name: 'privacy', component: () => import(/* webpackChunkName: "privacy" */ './views/Privacy.vue')},

    {path: '/404', alias: '*', name: 'notfound', component: () => import(/* webpackChunkName: "notfound" */ './views/NotFound.vue')},
  ],
});

router.beforeEach((to, from, next) => {
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    const token = localStorage.getItem(AUTH_TOKEN);

    if (token === null) {
      next({
        path: '/login',
        query: { next: to.fullPath },
      });
    } else {
      const decoded = jwtDecode(token);

      if (decoded.exp * 1000 < Date.now().valueOf()) {
        next({
          path: '/login',
          query: { next: to.fullPath },
        });
      } else {
        next();
      }
    }
  } else {
    next();
  }
});

export default router;
