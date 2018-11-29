const run = require('./function');
const path = require('path');

const {
  faaspath,
  uuid
} = require('../core/config');

module.exports = (handler) => (req, res) => {
  const data = {
    "path": req.url,
    "headers": req.headers || {},
    "httpMethod": req.method,
    "pathParameters": req.query || {},
    "body": req.body || {}
  }

  run(path.join(faaspath, uuid, handler), data, (result) => {
    res.status(result.statusCode || 200).send(result);
  })
}
