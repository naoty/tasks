const Dotenv = require("dotenv-webpack");

export default {
  apollo: {
    clientConfigs: {
      default: {
        httpEndpoint: process.env.GRAPHQL_ENDPOINT
      }
    }
  },
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
    "@nuxtjs/apollo",
    "@nuxtjs/vuetify"
  ]
}
