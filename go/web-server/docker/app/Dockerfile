FROM golang:alpine as base
# RUN apk update && apk upgrade && apk add --update alpine-sdk && apk add -no-cache git make cmake 

WORKDIR /go/src/going_in_circles/go/web-server/

COPY /go/src/going_in_circles/go/web-server/ .

RUN go mod download -x

ADD ./ .

RUN compile

ENTRYPOINT APP

wooknight/going_in_circles