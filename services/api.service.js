"use strict";

const ApiGateway = require("moleculer-web");

module.exports = {
  name: "api",
  mixins: [ApiGateway],

  // More info about settings: https://moleculer.services/docs/0.13/moleculer-web.html
  settings: {
    port: process.env.PORT || 3000,

    routes: [{
      path: "/api",
      whitelist: [
        // allow only lambda/** requests
        "**"
      ],
      aliases: {
        "lambda": "lambda.create",
        "DELETE lambda/:id": "lambda.delete"
      }
    }],

    // Serve assets from "public" folder
    assets: {
      folder: "public"
    }
  }
};