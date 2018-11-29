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
		},

		run: {
			params: {
				uuid: "string"
			},
			handler(ctx) {
				return "reuslt";
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