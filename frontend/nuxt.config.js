const Dotenv = require("dotenv-webpack");

export default {
  build: {
    extend(config, { }) {
      config.plugins.push(new Dotenv({ safe: true }));
    }
  },
  head: {
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" }
    ],
    title: "Tasks"
  },
  modules: [
    "@nuxtjs/vuetify"
  ]
}
