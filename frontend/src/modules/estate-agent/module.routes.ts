import { RouteConfig } from "vue-router";

const EstateAgentCreateUser = (resolve: any) => require(["@/modules/estate-agent/CreateUser.vue"], (m: any) => resolve(m.default));


const module: RouteConfig[] = [
  {
    path: '/create/estate-agent',
    component: EstateAgentCreateUser,
  }
];

export default module;
