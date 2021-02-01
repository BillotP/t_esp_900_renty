import { RouteConfig } from "vue-router";

const CreateEstateAgentUser = (resolve: any) => require(["@/modules/create-user/CreateEstateAgentUser.vue"], (m: any) => resolve(m.default));
const CreateTenantUser = (resolve: any) => require(["@/modules/create-user/CreateTenantUser.vue"], (m: any) => resolve(m.default));


const module: RouteConfig[] = [
  {
    path: '/create/estate-agent',
    component: CreateEstateAgentUser,
  },
  {
    path: '/create/tenant',
    component: CreateTenantUser,
  }
];

export default module;
