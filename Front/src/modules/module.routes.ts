import { RouteConfig } from "vue-router";

const TicketList = (resolve: any) => require(["@/components/TicketList.vue"], (m: any) => resolve(m.default));


const module: RouteConfig = {
    path: '/',
    component: TicketList,
};

export default module;
