import { RouteConfig } from "vue-router";
import AuthRoute from "@/modules/authentification/module.routes";

const TicketList = (resolve: any) => require(["@/modules/ticketList/TicketList.vue"], (m: any) => resolve(m.default));
const Disaster = (resolve: any) => require(["@/modules/ticketList/disaster/Disaster.vue"], (m: any) => resolve(m.default));

const module: RouteConfig[] = [
    {
        path: '/tickets',
        component: TicketList,
        children: [
            {
                path: "/dashboard",
                component: Disaster
            },
            {
                path: "/tenement",
                component: Disaster
            },
            {
                path: "/tickets",
                component: Disaster
            }

        ]
    },
    ...AuthRoute
];

export default module;
