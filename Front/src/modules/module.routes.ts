import { RouteConfig } from "vue-router";

const Board = (resolve: any) => require(["@/modules/board/Board.vue"], (m: any) => resolve(m.default));


const module: RouteConfig = {
    path: '/',
    component: Board,
};

export default module;
