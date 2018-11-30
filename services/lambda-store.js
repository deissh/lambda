"use strict";
const DbService = require("moleculer-db");

module.exports = {
	name: "db",
	mixins: [DbService],
	settings: {
		fields: ["_id", "name", "lang", "runtime", "functionVersion", "invokedFunctionArn", "sources"]
	},
	dependencies: [],
	actions: {},

	events: {},
	methods: {},
	created() {},
	started() {},
	stopped() {}
};