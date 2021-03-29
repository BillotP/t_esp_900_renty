import { RouteConfig } from "vue-router";
import { apolloClient } from '@/main';
import gql from "graphql-tag";


const TenantList = (resolve: any) => require(["@/modules/tenant/List.vue"], (m: any) => resolve(m.default));
const TenantProfile = (resolve: any) => require(["@/modules/tenant/Profile.vue"], (m: any) => resolve(m.default));
const TenantEdit = (resolve: any) => require(["@/modules/tenant/Edit.vue"], (m: any) => resolve(m.default));

const PROFILE_QUERY = gql`
  query profile {
    profile {
      ID
      __typename
    }
  }`;

const module: RouteConfig[] = [
  {
    path: '/tenants',
    component: TenantList,
  },
  {
    path: '/tenants/:id',
    component: TenantProfile,
  },
  {
    path: '/tenants/:id/edit',
    component: TenantEdit,
    beforeEnter: async (to, from, next) => {
      const res = await apolloClient.query({
        query: PROFILE_QUERY
      });

      const roles = {
        Company: 'companies',
        EstateAgent: 'estate-agents',
        Tenant: 'tenants'
      };

      if (from.path !== to.path && +to.params.id === +res.data.profile.ID && to.fullPath.indexOf(roles[res.data.profile.__typename]) > -1) {
        next();
        return;
      }

      next('/');
    }
  }
];

export default module;
