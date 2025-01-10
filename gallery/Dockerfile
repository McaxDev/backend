FROM golang:latest AS builder
ENV NAME=gallery
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /app/gallery
RUN go mod download
RUN go build -o gallery
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/gallery/gallery .
RUN apk update && apk add --no-cache libc6-compat tzdata
ENTRYPOINT [ "./gallery" ]
