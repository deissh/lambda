"use strict";

const service = {
	name: "lambda",
	settings: {

	},
	dependencies: [],	
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

module.exports = service;