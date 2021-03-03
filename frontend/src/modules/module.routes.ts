import { RouteConfig } from "vue-router";

import CreateUserRoute from "./create-user/module.routes";
import TenantRoute from "./tenant/module.routes";
import EstateAgent from "./estateAgent/module.routes";

const Login = (resolve: any) => require(["@/modules/authentication/Login.vue"], (m: any) => resolve(m.default));
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
    {
        path: '/login',
        component: Login
    },
    ...CreateUserRoute,
    ...TenantRoute
    ...EstateAgent
];

export default module;
