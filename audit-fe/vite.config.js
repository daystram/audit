import { defineConfig } from "vite";
import { createVuePlugin } from "vite-plugin-vue2";
import ViteComponents, { VuetifyResolver } from "vite-plugin-components";
import path from "path";

export default defineConfig({
  plugins: [
    createVuePlugin(),
    ViteComponents({
      globalComponentsDeclaration: true,
      customComponentResolvers: [VuetifyResolver()],
    }),
  ],
  server: {
    port: 8080,
  },
  resolve: {
    alias: [
      {
        find: "@",
        replacement: path.resolve(__dirname, "src"),
      },
    ],
  },
  build: {
    chunkSizeWarningLimit: 600,
    cssCodeSplit: false,
  },
});
