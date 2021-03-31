import Vue from "vue";

import Navbar from '@/components/Navbar.vue'
import TicketList from '@/components/TicketListComponent.vue'
import NavBarUnlogged from '@/components/NavbarUnlogged.vue'

Vue.component("navBar-field", Navbar);
Vue.component("navBarUnlogged-field", NavBarUnlogged)
Vue.component("ticketlist-field", TicketList);
