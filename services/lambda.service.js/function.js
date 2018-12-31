const hash = (Math.random().toString(16).slice(2) + Math.random().toString(16).slice(2) + Math.random().toString(16).slice(2)).substr(0, 32);

class Context {
  constructor(exitTimer, uuid) {
    this.exitTimer = exitTimer;

    this.awsRequestId = [
      hash.substr(0, 8), hash.substr(8, 4), hash.substr(12, 4), hash.substr(16, 4), hash.substr(20, 12)
    ].join("-");
    this.logStreamName = (new Date()).toISOString().substr(0, 10).replace(/-/g, "/") + "/[$LATEST]" + hash;
    this.memoryLimitInMB = "128";
    this.functionVersion = "$LATEST";
    this.functionUUID = uuid;
    this.invokedFunctionArn = "arn:aws:lambda:aws-region:1234567890123:function:" + uuid;
  }
  fail(err) {
    this._dumpError(err);
    this.exitTimer.unref();
  }
  done(error, message) {
    if (error) {
      this._dumpError(error);
      console.info(error);
    }
    this._dumpOutput(message);
    clearTimeout(this.exitTimer);
  }
  succeed(output) {
    this._dumpOutput(output);
    clearTimeout(this.exitTimer);
  }
  _dumpError(error) {
    console.info("ERROR");
    console.info("-".repeat(32));
    if (typeof (error) === "object") {
      if (error.constructor.name === "Error") {
        console.info(JSON.stringify({
          errorMessage: error.message ? error.message : "null",
          errorType: error.constructor.name,
          stackTrace: error.stack
        }, null, 4));
      } else {
        console.info(JSON.stringify(error, null, 4));
      }
    } else {
      console.info(JSON.stringify({
        errorMessage: error
      }, null, 4));
    }
  }

  _dumpOutput(output) {
    if (Object.prototype.toString.call(output) === "[object Object]") output = JSON.stringify(output, null, 4);
    console.info("OUTPUT");
    console.info("-".repeat(32));
    console.info(output);
  }
}

module.exports = (name, event, cb) => {
  const uuid = hash.substr(0, 8) + hash.substr(8, 16);
  // todo: rewrite this shit
  // context generator

  // main consoleic
  const exitTimer = setTimeout(() => {
    console.info("Lambda function " + uuid + " was timed out after 30 seconds");
    cb({});
  }, 30 * 1000);

  const context = new Context(exitTimer);

  const result = require(name).call({}, event, context, (error, output) => {
    if (typeof (error) !== "undefined" && error !== null) { // Normal callback call
      context.fail(error);
      cb({
        err: error,
        type: 2
      });
    }
    if (typeof (output) === "undefined") output = null;
    context.succeed(output);
    cb(output);
  });
  if (typeof (result) !== "undefined" && typeof (result.then) === "function") { // Promise returned or async function
    result
      .then((output) => {
        context.succeed(output);
        cb(output);
      })
      .catch((err) => {
        context.fail(err);
        cb({
          type: 3,
          err
        });
      });
  }

  clearTimeout(exitTimer);
};