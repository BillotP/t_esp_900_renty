import Vue from "vue";
import Router from "vue-router";
import { routes } from "./routes";

Vue.use(Router);

const router = new Router({
    mode: "history",
    routes
});

const unAuthPath = ["/login", "/register"];

router.beforeEach((to, _, next) => {
    var isLogged = localStorage.getItem("token") !== null ? true : false;

    if (!isLogged && !unAuthPath.includes(to.path)) {
        next({ path: unAuthPath[0] })
    }
    next();
})


export default router;
