FROM alpine:latest
WORKDIR /app
COPY geo .
RUN apk update && apk add --no-cache libc6-compat tzdata
ENTRYPOINT [ "./geo" ]
