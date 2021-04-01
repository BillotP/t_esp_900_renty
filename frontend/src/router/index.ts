import Vue from "vue";
import Router from "vue-router";
import { routes } from "./routes";

Vue.use(Router);

const router = new Router({
    mode: "history",
    routes
});

const unAuthPath = ["/login", "/register"];
const isLogged = localStorage.getItem("token") ? true : false;

router.beforeEach((to, _, next) => {
    if (!isLogged && !unAuthPath.includes(to.path)) {
        next({ path: unAuthPath[0] })
    }
    next();
})


export default router;
