FROM golang:alpine AS build

RUN apk add --no-cache git

WORKDIR /src

COPY . .

RUN go build -o /tmp/http-server ./cmd/http-server/main.go

FROM alpine:latest

COPY --from=build /tmp/http-server /sbin/http-server

EXPOSE 9090

CMD ["/sbin/http-server"]
