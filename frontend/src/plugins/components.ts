import Vue from "vue";
import Navbar from '@/components/Navbar.vue'
import NavBarUnlogged from '@/components/NavbarUnlogged.vue'
import {ModelObj} from 'vue-3d-model';
import VuePhoneNumberInput from 'vue-phone-number-input';
import 'vue-phone-number-input/dist/vue-phone-number-input.css';

Vue.component("navBar-field", Navbar);
Vue.component("navBarUnlogged-field", NavBarUnlogged);
Vue.component("model-obj", ModelObj);
Vue.component('vue-phone-number-input', VuePhoneNumberInput);
