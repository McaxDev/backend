FROM golang:latest AS builder
ENV NAME=wiki
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /app/wiki
RUN go mod download
RUN go build -o wiki
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/wiki/wiki .
RUN apk update && apk add --no-cache libc6-compat tzdata
ENTRYPOINT [ "./wiki" ]
