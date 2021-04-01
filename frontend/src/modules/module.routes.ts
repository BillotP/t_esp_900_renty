import { RouteConfig } from "vue-router";

import AuthRoute from "./authentication/module.routes";
import TenantRoute from "./tenant/module.routes";
import EstateAgentRoute from "./estate-agent/module.routes";
import PropertyRoute from "./property/module.routes";
import TicketRoute from "./ticket/module.routes";

const module: RouteConfig[] = [
    ...AuthRoute,
    ...TenantRoute,
    ...EstateAgentRoute,
    ...PropertyRoute,
    ...TicketRoute
];

export default module;
