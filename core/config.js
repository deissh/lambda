const requireEnv = (name) => {
  if (!process.env[name]) {
    throw new Error('You must set the ' + name + ' environment variable')
  }
  return process.env[name]
};

const config = {
  uuid: requireEnv('HANDLER_UUID'),
  faaspath: requireEnv('DEFAULT_PATH'),
  contextfile: process.env.DEFAULT_CONTEXT_FILE || null,
  yamlfilename: process.env.DEFAULT_YAML_FILENAME || 'serverless.yml',
  runner: {
    timeout: process.env.DEFAULT_TIMEOUT || 5
    // 60 sec
  },
}

module.exports = config;
