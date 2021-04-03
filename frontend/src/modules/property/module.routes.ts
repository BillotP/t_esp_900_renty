import { RouteConfig } from "vue-router";

const PropertyCreate = (resolve: any) => require(["@/modules/property/Create.vue"], (m: any) => resolve(m.default));
const PropertyList = (resolve: any) => require(["@/modules/property/List.vue"], (m: any) => resolve(m.default));
const PropertyProfile = (resolve: any) => require(["@/modules/property/Profile.vue"], (m: any) => resolve(m.default));


const module: RouteConfig[] = [
    {
        path: '/create/property',
        component: PropertyCreate,
    },
    {
        path: '/properties',
        component: PropertyList,
    },
    {
        path: '/property/:id',
        component: PropertyProfile,
    },
];

export default module;
