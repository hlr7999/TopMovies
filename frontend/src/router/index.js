import Vue from 'vue'
import Router from 'vue-router'
import HomeList from '@/components/HomeList'
import Description from '@/components/Description'
import axios from 'axios'

Vue.use(Router)
Vue.prototype.$axios = axios

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HomeList',
      component: HomeList
    },
    {
      path: '/description/:id',
      name: 'Description',
      component: Description
    }
  ]
})
