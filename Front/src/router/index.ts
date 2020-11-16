import Vue from "vue";
import Router from "vue-router";
import store from "../store/index";
import { routes } from "./routes";

Vue.use(Router);

const router = new Router({
  routes
});


router.beforeEach((to, from, next) => {
    next();
})

export default router;
