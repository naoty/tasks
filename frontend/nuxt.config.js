export default {
  env: {
    backendPort: process.env.BACKEND_PORT,
    clientBackendHost: process.env.CLIENT_BACKEND_HOST,
    serverBackendHost: process.env.SERVER_BACKEND_HOST
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
