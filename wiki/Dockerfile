FROM alpine:latest
RUN apk update && apk add --no-cache libc6-compat tzdata
WORKDIR /app
COPY app .
ENTRYPOINT [ "./app" ]
