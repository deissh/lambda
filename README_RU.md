# Simple Lambda Service

![img](https://blog.alexellis.io/content/images/2017/08/clip-1.png)

Lambda - это платформа для создания безсерверных функций с Docker. Любой процесс может быть упакован как функция и запущена.

## Начало

### Установка

#### Docker с Docker Compose

Если вы не хотите запускать локально проект то вам необходимо установить Docker и Docker Compose. После установки вам необходимо будет заполнить `docker-compose.yml`.
Внимание: необходимо создать простую функцию в папке указаной в `docker-compose.yml`

```bash
$ docker-compose up
```

or if you need rebuild

```bash
docker-compose up --build
```

#### Локально

Если у вас не установлен Docker и Docker Compose или вы не хотите уго устанавливать, то вам необходимо установить Node.JS и NPM (либо его аналоги).

Установка зависимостей

```bash
$ npm install
```

или с помошью `Yarn`

```bash
$ yarn
```

Так же необходимо установить и настроить Redis.

### Использование

#### Создание функции

Создайте новую папку для функции в `/temp` по умолчанию. Для примера, `123ws2x1`.

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

Пример файла `handler.js`

```bash
$ cat 123ws2x1/handler.js
module.exports = function (event, context, cb) {
  cb(null, {
    foo: 'bar',
    event,
    context
  });
};

```

Пример файла `serverless.yml`

```bash
$ cat 123ws2x1/serverless.yml
functions:
  - name: hello
    handler: handler.js
    events:
      - http
```

После создания функции необходимо перезапустить проект. Например с помощью Docker Compose

```bash
$ docker-compose restart
```

#### Использование функциии

Функция может быть вызвана с помощью HTTP запроса. 

```bash
$ curl http://127.0.0.1:3000/v1/{functionuuid}/{handlername}
```
Например

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
