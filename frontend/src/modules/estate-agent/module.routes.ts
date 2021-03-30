import { RouteConfig } from "vue-router";

const EstateAgentCreateUser = (resolve: any) => require(["@/modules/estate-agent/CreateUser.vue"], (m: any) => resolve(m.default));
const EstateAgentList = (resolve: any) => require(["@/modules/estate-agent/List.vue"], (m: any) => resolve(m.default));
const EstateAgentProfile = (resolve: any) => require(["@/modules/estate-agent/Profile.vue"], (m: any) => resolve(m.default));


const module: RouteConfig[] = [
  {
    path: '/create/estate-agent',
    component: EstateAgentCreateUser,
  },
  {
    path: '/estate-agents',
    component: EstateAgentList,
  },
  {
    path: '/estate-agent/:id',
    component: EstateAgentProfile,
  }
];

export default module;
