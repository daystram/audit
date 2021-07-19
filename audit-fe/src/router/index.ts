import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import { Dashboard, Home, Incidents, Monitor } from "@/views";

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
  {
    path: "/",
    component: Dashboard,
    children: [
      {
        path: "monitor",
        name: "dashboard:monitor",
        component: Monitor,
        meta: {
          title: "Monitor | Audit",
        },
      },
      {
        path: "incidents",
        name: "dashboard:incidents",
        component: Incidents,
        meta: {
          title: "Incidents | Audit",
        },
      },
    ],
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
