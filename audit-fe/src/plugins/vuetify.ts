import Vue from "vue";
import Vuetify from "vuetify/lib/framework";

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    dark: true,
    themes: {
      dark: {
        primary: "#4C1ED6",
        primaryDim: "#4721B8",
        secondary: "#D8315B",
      },
    },
    options: { customProperties: true },
  },
  icons: {
    iconfont: "mdi",
  },
});
