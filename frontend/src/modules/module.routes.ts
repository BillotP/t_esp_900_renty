import { RouteConfig } from "vue-router";

import AuthRoute from "./authentication/module.routes";
import TenantRoute from "./tenant/module.routes";
import EstateAgentRoute from "./estate-agent/module.routes";
import PropertyRoute from "./property/module.routes";
import TicketRoute from "./ticket/module.routes";
import CompanyRoute from "./company/module.route";

const module: RouteConfig[] = [
    ...AuthRoute,
    ...TenantRoute,
    ...CompanyRoute,
    ...EstateAgentRoute,
    ...PropertyRoute,
    ...TicketRoute,
    {
        path: "/**",
        redirect: "/dashboard"
    }
];

export default module;
