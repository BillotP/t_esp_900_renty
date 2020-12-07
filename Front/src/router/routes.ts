/*import Vue from "vue";
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
*/

// import Login from "@/modules/authentification/Login.vue";

const Root = (resolve: any) => require(['@/layout/Root.vue'], (m: any) => resolve(m.default));

import modules from "@/modules/module.routes";

import { RouteConfig } from 'vue-router';

export const routes: RouteConfig[] = [
    {
        path: "/",
        component: Root,
        children: [
            ...modules,
        ]
    }
];
