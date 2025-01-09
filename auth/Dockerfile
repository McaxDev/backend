FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /app/auth
RUN go mod download
RUN go build -o auth
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/auth/auth .
RUN apk update && apk add --no-cache libc6-compat
ENTRYPOINT [ "./auth" ]
