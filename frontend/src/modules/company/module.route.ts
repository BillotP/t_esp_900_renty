import { RouteConfig } from "vue-router";

const Dashboard = (resolve: any) => require(["@/modules/dashboard/Dashboard.vue"], (m: any) => resolve(m.default));

const EditCompany = (resolve: any) => require(["@/modules/company/EditCompany.vue"], (m: any) => resolve(m.default));

const module: RouteConfig[] = [
    {
        path: '/companies/:id/edit',
        component: EditCompany,
    },
    {
        path: '/dashboard',
        component: Dashboard
    },
];

export default module;
