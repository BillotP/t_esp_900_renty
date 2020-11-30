import { RouteConfig } from "vue-router";

const AuthClient = (resolve: any) => require(["@/modules/authentification/loginAsCustomer/Auth.vue"], (m: any) => resolve(m.default));
const AuthCompany = (resolve: any) => require(["@/modules/authentification/loginAsCompany/Auth.vue"], (m: any) => resolve(m.default));


const module: RouteConfig[] = [
  {
    path: '/login/client',
    component: AuthClient,
  },
  {
    path: '/login/company',
    component: AuthCompany,
  }
];

export default module;
