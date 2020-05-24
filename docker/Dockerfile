FROM golang:latest

RUN mkdir /go/src/sendgrid

WORKDIR /go/src/sendgrid

RUN go get github.com/sendgrid/sendgrid-go

ADD . /go/src/sendgrid
