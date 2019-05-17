import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

Vue.config.productionTip = false

/*import Vuikit from 'vuikit'
import VuikitIcons from '@vuikit/icons'

import '@vuikit/theme'

Vue.use(Vuikit)
Vue.use(VuikitIcons)*/

import VueResource from 'vue-resource'
Vue.use(VueResource)

import BootstrapVue from 'bootstrap-vue'
Vue.use(BootstrapVue)

import 'bootstrap/dist/css/bootstrap.css'
import './assets/skins/dark.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import VueMaterialIcon from 'vue-material-icon'
Vue.component('mat-icon', VueMaterialIcon)

import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(far)
library.add(fas)

Vue.component('fa-icon', FontAwesomeIcon)

import './assets/style.css'

import Navigation from './navigation'

Vue.use(Navigation)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
