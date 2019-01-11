FROM golang:1.10 AS builder

WORKDIR $GOPATH/src/github.com/deissh/lambda
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /www ./cmd/www

FROM scratch
ENV GIN_MODE=release
COPY --from=builder /www ./www
EXPOSE 3000
ENTRYPOINT ["./www"]