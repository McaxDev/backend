FROM golang:latest AS builder
ENV NAME=games
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /app/games
RUN go mod download
RUN go build -o games
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/games/games .
RUN apk update && apk add --no-cache libc6-compat tzdata
ENTRYPOINT [ "./games" ]
