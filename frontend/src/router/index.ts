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
    let isLogged = localStorage.getItem("token") !== null ? true : false;
    let isExpired = true

    if (isLogged) {
        let exp = localStorage.getItem("exp")
        isExpired = Date.now() >= Number(exp) * 1000 ? true : false;
    }

    if ((!isLogged || isExpired) && !unAuthPath.includes(to.path)) {
        next({ path: unAuthPath[0] })
    }
    next();
})


export default router;
