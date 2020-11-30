import { RouteConfig } from "vue-router";
import AuthRoute from "./authentification/module.routes";

const TicketList = (resolve: any) => require(["@/components/TicketList.vue"], (m: any) => resolve(m.default));


const module: RouteConfig[] = [
    {
        path: '/',
        component: TicketList,
    },
    ...AuthRoute
];

export default module;
