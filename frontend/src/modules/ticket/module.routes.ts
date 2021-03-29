import { RouteConfig } from "vue-router";
import { apolloClient } from '@/main';
import gql from "graphql-tag";

const TicketCreate = (resolve: any) => require(["@/modules/ticket/Create.vue"], (m: any) => resolve(m.default));

const module: RouteConfig[] = [
  {
    path: '/create/ticket',
    component: TicketCreate,
  },
];

export default module;
