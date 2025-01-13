FROM golang:latest AS builder
ENV NAME=guild
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /app/guild
RUN go mod download
RUN go build -o guild
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/guild/guild .
RUN apk update && apk add --no-cache libc6-compat tzdata
ENTRYPOINT [ "./guild" ]
