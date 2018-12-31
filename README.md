# Simple Lambda Service

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/c4d781bc8ecd430b95385f8a4bc0fec5)](https://www.codacy.com/app/Deissh/lambda?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=deissh/lambda&amp;utm_campaign=Badge_Grade)
[![David](https://img.shields.io/david/deissh/lambda.svg)](https://david-dm.org/deissh/lambda)
[![Known Vulnerabilities](https://snyk.io/test/github/deissh/lambda/badge.svg)](https://snyk.io/test/github/deissh/lambda)

[![GitHub issues](https://img.shields.io/github/issues/deissh/lambda.svg)](https://github.com/deissh/lambda/issues)
[![GitHub stars](https://img.shields.io/github/stars/deissh/lambda.svg)](https://github.com/deissh/lambda/stargazers)
[![Twitter](https://img.shields.io/twitter/url/https/github.com/deissh/lambda.svg?style=social)](https://twitter.com/intent/tweet?text=Wow:&url=https%3A%2F%2Fgithub.com%2Fdeissh%2Flambda)

![img](https://blog.alexellis.io/content/images/2017/08/clip-1.png)

[RU](https://github.com/deissh/lambda/blob/master/README_RU.md)

Lambda is a framework for building Serverless functions with Docker which has first-class support for metrics. Any process can be packaged as a function enabling you to consume a range of web events without repetitive boiler-plate coding.

## Get starting

### Deployment

``` bash
# Install dependencies
npm install

# Start developing with REPL
npm run dev

# Start production
npm start

# Run unit tests
npm test

# Run continuous test mode
npm run ci

# Run ESLint
npm run lint
```

#### Docker with Docker Compose

If you won't use Node.JS and NVM in your workstation need to install Docker and also Docker-Compose. Then you have done you can configure `docker-composer.yml` and run it.
Warning: You must create a simple function in the temp folder (or inside a folder is chosen).

```bash
$ docker-compose up
```

or if you need rebuild

```bash
docker-compose up --build
```

#### Node.JS in workstation

If you don't have Docker or won't install it, you can install Runner locally with Node.JS. Install Node.JS locally if you don't have it by NVM or analogues.

```bash
$ npm install
```

or with `Yarn`

```bash
$ yarn
```

Also, you need to install and configure Redis locally or remotely.

### Using

#### Write your function

Create a new folder for your work, the default is `/temp` folder. For example, `123ws2x1` is function UUID.

```bash
(in project folder)
$ mkdir temp && cd temp
$ mkdir 123ws2x1
$ touch 123ws2x1/handler.js
$ touch 123ws2x1/serverless.yml
```


```bash
$ tree .
.
└── 123ws2x1
    ├── handler.js - Contain your code.
    └── serverless.yml - Contain configurations
```

Code example with `serverless.yml`

```bash
$ cat 123ws2x1/handler.js
module.exports = function (event, context, cb) {
  cb(null, {
    foo: 'bar',
    event,
    context
  });
};

$ cat 123ws2x1/serverless.yml
functions:
  - name: hello
    handler: handler.js
    events:
      - http
```

When you have done, you need to restart Docker-Compose or locally project.

```bash
$ docker-compose restart
```

#### Trigger you function

Your function can be triggered by HTTP request, for example, you can use cURL or others HTTP clients.

```bash
$ curl http://127.0.0.1:3000/v1/{functionuuid}/{handlername}
```

for example

```bash
$ curl http://127.0.0.1:3000/v1/123ws2x1/hello

{
  "foo": "bar",
  "event": {
    "path": "/123ws2x1/hello",
    "headers": {
      "host": "127.0.0.1:3000",
      "connection": "keep-alive",
      "cache-control": "max-age=0",
      "upgrade-insecure-requests": "1",
     },
    "httpMethod": "GET",
    "pathParameters": {
      
    },
    "body": {
      
    }
  },
  "context": {
    "awsRequestId": "6654c38a-5e1e-6031-59f8-c41ffb8f65dc",
    "logStreamName": "2018/11/23/[$LATEST]6654c38a5e1e603159f8c41ffb8f65dc",
    "memoryLimitInMB": "128",
    "functionVersion": "$LATEST",
    "functionUUID": "123ws2x1",
    "invokedFunctionArn": "arn:aws:lambda:aws-region:1234567890123:function:123ws2x1"
  }
}
```
