import { RouteConfig } from "vue-router";

const TenantList = (resolve: any) => require(["@/modules/tenant/List.vue"], (m: any) => resolve(m.default));
const TenantProfile = (resolve: any) => require(["@/modules/tenant/Profile.vue"], (m: any) => resolve(m.default));

const module: RouteConfig[] = [
  {
    path: '/tenants',
    component: TenantList,
  },
  {
    path: '/tenants/:id',
    component: TenantProfile,
  }
];

export default module;
