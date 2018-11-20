FROM node:10-alpine
ENV NODE_ENV production
MAINTAINER FunctionRunner deissh@yandex.ru

WORKDIR /usr/src/app
ADD package.json package.json
RUN npm install pino-elasticsearch -g
RUN npm install --production
RUN rm -f .npmrc
ADD . .
CMD ["npm", "start"]
