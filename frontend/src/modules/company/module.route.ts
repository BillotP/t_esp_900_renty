import { RouteConfig } from "vue-router";

const EditCompany = (resolve: any) => require(["@/modules/company/EditCompany.vue"], (m: any) => resolve(m.default));

const module: RouteConfig[] = [
    {
        path: '/companies/:id/edit',
        component: EditCompany,
    },
];

export default module;
