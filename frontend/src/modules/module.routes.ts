import { RouteConfig } from "vue-router";

import TenantRoute from "./tenant/module.routes";
import EstateAgentRoute from "./estate-agent/module.routes";
import PropertyRoute from "./property/module.routes";
import TicketRoute from "./ticket/module.routes";

const Login = (resolve: any) => require(["@/modules/authentication/Login.vue"], (m: any) => resolve(m.default));

const module: RouteConfig[] = [
    {
        path: '/login',
        component: Login
    },
    ...TenantRoute,
    ...EstateAgentRoute,
    ...PropertyRoute,
    ...TicketRoute
];

export default module;
