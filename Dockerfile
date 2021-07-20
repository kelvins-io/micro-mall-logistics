FROM golang:1.13.15-buster as gobuild
WORKDIR $GOPATH/src/gitee.com/cristiane/micro-mall-logistics
COPY . .
ENV GOPROXY=https://goproxy.io,direct
RUN bash ./build.sh
# FROM alpine:latest as gorun
FROM ubuntu:latest as gorun
WORKDIR /www/
COPY --from=gobuild /go/src/gitee.com/cristiane/micro-mall-logistics/micro-mall-logistics .
COPY --from=gobuild /go/src/gitee.com/cristiane/micro-mall-logistics/etc ./etc
CMD ["./micro-mall-logistics"]
