## Build Docker Image like this:
# docker build -f docker/Dockerfile -t backend .

FROM alpine:3.19
WORKDIR /app
RUN apk add --no-cache libc6-compat
COPY ../build/linux/weather-proxy-api-linux ./

EXPOSE 3000

CMD ["/app/weather-api-linux"]