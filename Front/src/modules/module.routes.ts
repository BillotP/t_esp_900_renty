import { RouteConfig } from "vue-router";

const Disaster = (resolve: any) => require(["@/modules/disaster/Disaster.vue"], (m: any) => resolve(m.default));


const module: RouteConfig = {
    path: '/disaster',
    component: Disaster,
};

export default module;
