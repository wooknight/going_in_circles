FROM golang:alpine as base
# RUN apk update && apk upgrade && apk add --update alpine-sdk && apk add -no-cache git make cmake 

WORKDIR /go/src/GoingInCircles/Go/web-server/

COPY /go/src/GoingInCircles/Go/web-server/ .

RUN go mod download -x

ADD ./ .

RUN compile

ENTRYPOINT APP

