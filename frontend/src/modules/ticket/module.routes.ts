import { RouteConfig } from "vue-router";
import { apolloClient } from '@/main';
import gql from "graphql-tag";

const TicketCreate = (resolve: any) => require(["@/modules/ticket/Create.vue"], (m: any) => resolve(m.default));
const TicketList = (resolve: any) => require(["@/modules/ticket/List.vue"], (m: any) => resolve(m.default));
const TicketProfile = (resolve: any) => require(["@/modules/ticket/Profile.vue"], (m: any) => resolve(m.default));

const module: RouteConfig[] = [
  {
    path: '/create/ticket',
    component: TicketCreate,
  },
  {
    path: '/tickets',
    component: TicketList,
  },
  {
    path: '/ticket/:id',
    component: TicketProfile,
  },
];

export default module;
