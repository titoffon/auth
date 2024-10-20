FROM golang:1.23.2-alpine3.20 AS builder

COPY . /github.com/titoffon/auth/
WORKDIR /github.com/titoffon/auth/

RUN go mod download
RUN go build -o ./bin/auth cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/titoffon/auth/bin/auth .

CMD ["./auth"]