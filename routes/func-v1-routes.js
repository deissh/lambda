/**
 * @name func-v1-api
 * @description This module packages the Func API.
 */
'use strict';

const yaml = require('js-yaml');
const fs = require('fs');
const path = require('path');

const hydraExpress = require('hydra-express');
const hydra = hydraExpress.getHydra();
const express = hydraExpress.getExpress();
const jwtAuth = require('fwsp-jwt-auth');
const ServerResponse = require('fwsp-server-response');

const {
  faaspath,
  uuid,
  yamlfilename,
} = require('../core/config');

const handler = require('../core/handler');

let yamlfile = yaml.safeLoad(
  fs.readFileSync(
    path.join(faaspath, uuid, yamlfilename),
    'utf8'
  )
);

let serverResponse = new ServerResponse();

serverResponse.enableCORS(true);
express.response.sendError = function (err) {
  serverResponse.sendServerError(this, {
    result: {
      error: err
    }
  });
};
express.response.sendOk = function (result) {
  serverResponse.sendOk(this, {
    result
  });
};

let api = express.Router();

yamlfile.functions.forEach(func => {
  console.info('add new function ' + func.name)

  func.events.forEach(event => {
    console.info('add event trigger by ' + event)
    if (event === 'http') {
      api.all(`/${uuid}/${func.name}`, handler(func.handler));
    } else if (event === 'all') {
      // nc.subscribe(`event.*.*.*`, handler(func.handler))
    }
  })

})

module.exports = api;
