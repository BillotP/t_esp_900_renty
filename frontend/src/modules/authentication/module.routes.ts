import { RouteConfig } from "vue-router";

const Login = (resolve: any) => require(["@/modules/authentication/Login.vue"], (m: any) => resolve(m.default));
const Register = (resolve: any) => require(["@/modules/authentication/Register.vue"], (m: any) => resolve(m.default));

const module: RouteConfig[] = [
  {
    path: '/login',
    component: Login,
  },
  {
    path: '/register',
    component: Register,
  }
];

export default module;
