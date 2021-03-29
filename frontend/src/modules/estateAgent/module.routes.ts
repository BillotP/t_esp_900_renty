import { RouteConfig } from "vue-router";

const CreateProperty = (resolve: any) => require(["@/modules/estateAgent/property/CreateProperty.vue"], (m: any) => resolve(m.default));
const CreateTicket = (resolve: any) => require(["@/modules/estateAgent/property/CreateTicket.vue"], (m: any) => resolve(m.default));


const module: RouteConfig[] = [
  {
    path: '/estateAgent/createProperty',
    component: CreateProperty,
  },
  {
    path: '/tenant/:id/CreateTicket/',
    component: CreateTicket
  }
];

export default module;
