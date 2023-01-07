FROM golang:1.18.6-alpine AS build

WORKDIR /app

ADD internal /app/internal
ADD pkg /app/pkg
ADD main.go /app
ADD go.mod /app
ADD go.sum /app
ADD bootstrap /app/bootstrap

RUN CGO_ENABLED=0 GOOS=linux
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o server .


FROM alpine:latest

WORKDIR /app

COPY --from=build /app/server .

EXPOSE 9001
ENV TZ=Asia/shanghai

ENTRYPOINT ["/app/server"]