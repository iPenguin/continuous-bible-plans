FROM golang:1.16-buster

COPY build/cbp /go/app/
#VOLUME /go/app/

ENTRYPOINT ["/go/app/cbp"]


