import Vue from 'vue'
import App from './App'
import router from './router'
import Vuex from 'vuex'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

Vue.config.productionTip = false

Vue.use(Vuex)
Vue.use(ElementUI)

const store = new Vuex.Store({
  state:{
    currentPage: 1,
    sortWay: ""
  },
  mutations:{
    newPage(state, page) {
      state.currentPage = page
    },
    newSort(state, sort) {
      state.sortWay = sort
    }
  }
})

new Vue({
  el: '#app',
  render: h => h(App),
  router,
  store
})
