import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  routes: [
    {
      name: 'Home',
      path: '/home',
      component: () => import(/* webpackChunkName: "about" */ './views/Home.vue')
    },
    {
      path: '/library',
      component: () => import(/* webpackChunkName: "about" */ './views/library/View.vue'),
      children: [
        {
          name: 'Explore',
          path: 'explore',
          component: () => import(/* webpackChunkName: "about" */ './views/library/Explore.vue')
        }, {
          name: 'Collections',
          path: 'collections',
          component: () => import(/* webpackChunkName: "about" */ './views/library/Collections.vue')
        }, {
          name: 'Available',
          path: 'available',
          component: () => import(/* webpackChunkName: "about" */ './views/library/Available.vue')
        }, {
          name: 'All Games',
          path: 'all',
          component: () => import(/* webpackChunkName: "about" */ './views/library/All.vue')
        }
      ]
    },
    {
      name: 'News',
      path: '/news',
      component: () => import(/* webpackChunkName: "about" */ './views/News.vue')
    },
    {
      name: 'Store',
      path: '/store',
      component: () => import(/* webpackChunkName: "about" */ './views/Store.vue')
    },
    {
      name: 'Settings',
      path: '/settings',
      component: () => import(/* webpackChunkName: "about" */ './views/Settings.vue')
    },
    {
      name: 'User',
      path: '/user',
      component: () => import(/* webpackChunkName: "about" */ './views/User.vue')
    }
  ]
})
