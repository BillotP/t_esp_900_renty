import { RouteConfig } from "vue-router";

const CreateProperty = (resolve: any) => require(["@/modules/estateAgent/property/CreateProperty.vue"], (m: any) => resolve(m.default));


const module: RouteConfig[] = [
  {
    path: '/estateAgent/createProperty',
    component: CreateProperty,
  }
];

export default module;
