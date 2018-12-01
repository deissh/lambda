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
    hello(ctx) {
      return {
        data: "I'm alive!",
        status: 200
      };
    },

    /**
     * Welcome a username
     *
     * @param {String} name - User name
     */
    create: {
      params: {

      },
      handler(ctx) {
        
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