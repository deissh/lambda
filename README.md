# func-service

![img](https://blog.alexellis.io/content/images/2017/08/clip-1.png)

Jinesa FaaS is a framework for building Serverless functions with Docker which has first-class support for metrics. Any process can be packaged as a function enabling you to consume a range of web events without repetitive boiler-plate coding.

## Get starting

### Deployment

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
