FROM golang:1.23.0-alpine AS builder
RUN apk add --no-cache gcc musl-dev libwebp-dev
WORKDIR /app
COPY . .
RUN CGO_ENABLED=1 ./build-static.sh

ENV GOMEMLIMIT=2GiB
ENV GOMAXPROCS=1

FROM alpine:latest
RUN apk add --no-cache bash ffmpeg psmisc
WORKDIR /app
COPY --from=builder /app/teamgramd/ /app/
RUN chmod +x /app/docker/entrypoint.sh
ENTRYPOINT /app/docker/entrypoint.sh
