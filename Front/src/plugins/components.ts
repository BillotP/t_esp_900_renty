import Vue from "vue";

/*
example :
import FormFieldComponent from "@/components/forms/layout/Field.vue";

Vue.component("form-field", FormFieldComponent);
*/

import Navbar from '@/components/Navbar.vue'
import TicketList from '@/components/TicketListComponent.vue'

Vue.component("navBar-field", Navbar);
Vue.component("ticketlist-field", TicketList);
