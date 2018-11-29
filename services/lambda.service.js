"use strict";

module.exports = {
	name: "lambda",
	settings: {

	},
	dependencies: [],	
	actions: {
		hello() {
			return "I'm alive!";
		},

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