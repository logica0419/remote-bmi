import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react({
      babel: {
        plugins: [
          [
            "@babel/plugin-transform-react-jsx",
            { runtime: "automatic", importSource: "@emotion/react" },
            "emotion-loader",
          ],
        ],
      },
    }),
  ],
  server: {
    port: 3001,
    proxy: {
      "/api": {
        target: "http://localhost:3000",
        changeOrigin: true,
      },
    },
  },
  build: {
    outDir: "../server/router/static",
  },
});
