FROM golang:1.10 AS builder

WORKDIR $GOPATH/src/github.com/deissh/lambda
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /manager ./cmd/manager

FROM scratch
COPY --from=builder /manager ./manager
EXPOSE 3000
ENTRYPOINT ["./manager"]