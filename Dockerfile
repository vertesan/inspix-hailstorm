# syntax=docker/dockerfile:1

FROM golang:1.22.1 AS build-stage
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./hailstorm

FROM alpine:latest AS build-release-stage
# RUN echo -e 'busybox ls -alFh --color' > /usr/bin/ll && \
#   chmod +x /usr/bin/ll
RUN apk add --no-cache bash openssh git
WORKDIR /app
RUN mkdir cache && mkdir masterdata && mkdir secrets
COPY --from=build-stage /app/hailstorm ./hailstorm
COPY ./docker-entrypoint.sh ./entrypoint.sh
RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
