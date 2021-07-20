import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import { Dashboard, Home, Incidents, Manage, Monitor } from "@/views";
import { authenticatedOnly, callback, login, logout, unAuthenticatedOnly } from "@/auth";

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
      {
        path: "manage",
        name: "dashboard:manage",
        component: Manage,
        meta: {
          title: "Manage | Audit",
        },
      },
    ],
  },
  {
    path: "/login",
    name: "login",
    beforeEnter: unAuthenticatedOnly,
    component: login,
  },
  {
    path: "/logout",
    name: "logout",
    beforeEnter: authenticatedOnly,
    component: logout,
  },
  {
    path: "/callback",
    name: "callback",
    beforeEnter: unAuthenticatedOnly,
    component: callback,
  },
  {
    path: "*",
    redirect: { name: "home" },
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
