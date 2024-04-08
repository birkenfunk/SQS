## Build Docker Image like this:
# docker build -f docker/Dockerfile -t backend .

FROM alpine:3.19
WORKDIR /app
RUN apk add --no-cache libc6-compat
COPY ../build/linux/weather-proxy-api-linux ./

EXPOSE 3000

RUN adduser -D server
USER server

CMD ["/app/weather-proxy-api-linux"]