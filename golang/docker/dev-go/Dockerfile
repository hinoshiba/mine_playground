FROM ubuntu:20.04

ENV GOVERSION 1.16.3

ENV GOROOT /opt/go
ENV GOPATH /go/src

ENV LANG en_US.UTF-8
#ENV http_proxy http://proxy.example.com:8080/
#ENV https_proxy http://proxy.example.com:8080/
ENV http_proxy http://proxy.i.hinoshiba.com:8080/
ENV https_proxy http://proxy.i.hinoshiba.com:8080/
ENV DEBIAN_FRONTEND=noninteractive

RUN apt update && \
    apt install -y git curl wget tzdata make gcc && \
## feny
    apt install -y libx11-dev libxcursor-dev libgl1-mesa-dev xorg-dev gcc-multilib gcc-mingw-w64 binutils-mingw-w64 && \
## feny-done
## qt
    apt install -y build-essential libglu1-mesa-dev libpulse-dev libglib2.0-dev && \
## qt-done
    apt clean && \
    rm -rf /var/lib/apt/lists/* && \
    ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

RUN cd /opt && wget -nv --show-progress --progress=bar:force:noscroll https://golang.org/dl/go${GOVERSION}.linux-amd64.tar.gz && \
    tar zxf go${GOVERSION}.linux-amd64.tar.gz && rm go${GOVERSION}.linux-amd64.tar.gz && \
    ln -s /opt/go/bin/go /usr/bin/ && \
## qt
    export GO111MODULE=off; go get -v github.com/therecipe/qt/cmd/... && $(go env GOPATH)/bin/qtsetup test && $(go env GOPATH)/bin/qtsetup -test=false && \
## qt-done
    mkdir -p ${GOPATH} && \
    mkdir /.cache && \
    chmod 777 /.cache

WORKDIR ${GOPATH}
