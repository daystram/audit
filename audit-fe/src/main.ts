import Vue from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
import vuetify from "./plugins/vuetify";
import { StatusMixin } from "./constants/status";

import "@/styles/App.sass";

Vue.config.productionTip = false;
Vue.mixin(StatusMixin);
new Vue({
  router,
  vuetify,
  render: (h) => h(App),
}).$mount("#app");
