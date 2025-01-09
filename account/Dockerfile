FROM golang:latest AS builder
ENV NAME=account
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /app/account
RUN go mod download
RUN go build -o account
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/account/account .
RUN apk update && apk add --no-cache libc6-compat tzdata
ENTRYPOINT [ "./account" ]
