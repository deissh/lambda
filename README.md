# Simple Lambda Service

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/c4d781bc8ecd430b95385f8a4bc0fec5)](https://www.codacy.com/app/Deissh/lambda?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=deissh/lambda&amp;utm_campaign=Badge_Grade)

[![GitHub issues](https://img.shields.io/github/issues/deissh/lambda.svg)](https://github.com/deissh/lambda/issues)
[![GitHub stars](https://img.shields.io/github/stars/deissh/lambda.svg)](https://github.com/deissh/lambda/stargazers)
[![Twitter](https://img.shields.io/twitter/url/https/github.com/deissh/lambda.svg?style=social)](https://twitter.com/intent/tweet?text=Wow:&url=https%3A%2F%2Fgithub.com%2Fdeissh%2Flambda)

![img](https://blog.alexellis.io/content/images/2017/08/clip-1.png)

Lambda is a framework for building Serverless functions with Docker which has first-class support for metrics. Any process can be packaged as a function enabling you to consume a range of web events without repetitive boiler-plate coding.

## API

Your function available on custom port. For example `8080` is function port.

### Triger function

Allow all http methods

```bash
curl -X GET \
  http://localhost:3000/v1/function/{uuid}
```

### Create new function

```bash
curl -X POST \
  http://localhost:3000/v1/create \
  -H 'Content-Type: application/json' \
  -d '{
	"name": "example",
	"runtime": {
		"executor": "/bin/cat",
		"cmd": ""
	},
	"repository": {
		"image": "deissh/lambda-runner:latest"
	},
	"service": {
		"port": "8080"
	}
}'
```

Result

```json
{
    "uuid": "c80d737d0040dff9c9b0341908273dfd71f66f4e5eb8302a0b9d8b26e9b87089"
}
```

### Inspect function

`/v1/inspect/:uuid`

```bash
curl -X GET \
  http://localhost:3000/v1/inspect/{uuid}
```

### Delete function

`/v1/delete/:uuid`

You need change uuid befour use.

```bash
curl -X GET \
  http://localhost:3000/v1/delete/{uuid}
```
