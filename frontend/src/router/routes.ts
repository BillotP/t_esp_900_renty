const Root = (resolve: any) => require(['@/layout/Root.vue'], (m: any) => resolve(m.default));

import modules from "@/modules/module.routes";

import { RouteConfig } from 'vue-router';

export const routes: RouteConfig[] = [
  {
    path: "/",
    component: Root,
    children: [
      ...modules,
    ]
  },
];
