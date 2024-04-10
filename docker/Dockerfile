## Build Docker Image like this:
# docker build -f docker/Dockerfile -t backend .

FROM alpine:3.19
WORKDIR /app
RUN apk add --no-cache musl-obstack=1.2.3-r2 libucontext=1.2-r2 gcompat=1.1.0-r4
COPY ../build/linux/weather-proxy-api-linux ./

EXPOSE 3000

RUN adduser -D server
USER server

CMD ["/app/weather-proxy-api-linux"]