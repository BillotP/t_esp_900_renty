import Vue from "vue";
import Router from "vue-router";
import store from "../store/index";
import {routes} from "./routes";

Vue.use(Router);

const router = new Router({
    mode: "history",
    routes
});

router.beforeEach((to, from, next) => {
    next();
    if (to.path !== '/login' && !localStorage.getItem('username')) next({path: '/login'})
    else next();
})


export default router;
