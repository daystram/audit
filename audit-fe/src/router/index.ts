import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import { Home } from "@/views";

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "home",
    component: Home,
    meta: {
      title: "Audit",
    },
  },
];

const router = new VueRouter({
  mode: "history",
  base: import.meta.env.BASE_URL,
  routes,
});

router.beforeEach((to, from, next) => {
  document.title = to?.meta?.title || "Audit";
  next();
});

export default router;
