FROM golang:1.15

ENV PB_VER 3.10.1

ENV PB_URL https://github.com/protocolbuffers/protobuf/releases/download/v${PB_VER}/protoc-${PB_VER}-linux-x86_32.zip
RUN apt-get -qq update && apt-get -qqy install curl git make unzip gettext  rsync

RUN mkdir -p /tmp/protoc &&  \
    curl -L ${PB_URL} > /tmp/protoc/protoc.zip && \
    cd /tmp/protoc && \
    unzip protoc.zip && \
    cp /tmp/protoc/bin/protoc /usr/local/bin && \
    cp -R /tmp/protoc/include/* /usr/local/include && \
    chmod go+rx /usr/local/bin/protoc && \
    cd /tmp && \
    rm -r /tmp/protoc

#GetthesourcefromGitHub
RUN go get -u google.golang.org/grpc

#Installprotoc-gen-go
RUN go get -u github.com/golang/protobuf/protoc-gen-go

##Install protoc-gen -twirp 
RUN go get github.com/twitchtv/twirp/protoc-gen-twirp
