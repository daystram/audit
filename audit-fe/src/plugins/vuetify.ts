import Vue from "vue";
import Vuetify from "vuetify/lib/framework";

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    dark: true,
    themes: {
      dark: {
        primary: "#00ADFF",
        primaryDim: "#0082BF",
        secondary: "#D8315B",
      },
    },
    options: { customProperties: true },
  },
  icons: {
    iconfont: "mdi",
  },
});
