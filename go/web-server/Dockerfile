FROM golang:1.16-alpine

# RUN apk update && apk upgrade && apk add --update alpine-sdk && apk add -no-cache git make cmake 

RUN mkdir /app

WORKDIR /go/src/going_in_circles/go/web-server

RUN pwd
ENV GO111MODULE on
COPY . .
RUN go mod download -x


ADD ./ .


RUN go build -o /web-server

EXPOSE 8000

CMD ["/web-server"]
