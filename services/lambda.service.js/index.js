"use strict";

const service = {
  name: "lambda",
  settings: {

  },
  dependencies: [
    "lambda-store"
    // "lambda-drive"
  ],
  actions: {
    /**
     * Welcome a username
     *
     * @param {String} name - User name
     */
    welcome: {
      params: {
        name: "string"
      },
      handler(ctx) {
        return `Welcome, ${ctx.params.name}`;
      }
    },
    delete: {
      params: {
        id: "string"
      },
      handler(ctx) {

      }
    }
  },
  events: {

  },
  methods: {

  },
  created() {

  },
  started() {

  },
  stopped() {

  }
};

module.exports = service;