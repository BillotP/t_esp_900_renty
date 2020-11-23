import Vue from "vue";
import Router from "vue-router";
import Home from '../views/Home.vue'
import Dashboard from '../components/Dashboard.vue'
import Plannings from '../components/Plannings.vue'
import TicketList from '../components/TicketList.vue'

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/dashboard',
      name: 'Dashboard',
      component: Dashboard
    },
    {
      path: '/plannings',
      name: 'Plannings',
      component: Plannings
    },
    {
      path: '/tickets',
      name: 'TicketList',
      component: TicketList
    }
  ]
})
