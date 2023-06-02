FROM alpine:latest
RUN apk add --no-cache tzdata && cp /usr/share/zoneinfo/PRC /etc/localtime && echo "PRC" > /etc/timezone

WORKDIR /app

COPY ./assets /app/assets
COPY ./configs /app/configs
COPY ./templates /app/templates
COPY ./.env /app/.env
COPY ./server /app/server

EXPOSE 8080

CMD ["/app/server"]
