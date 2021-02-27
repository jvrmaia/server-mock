FROM golang:1.14.1-alpine AS build-env

ADD . /src

RUN cd /src && go build -o server-mock

FROM alpine:3

WORKDIR /app
COPY --from=build-env /src/server-mock /app/

EXPOSE 8080

ENTRYPOINT ./server-mock