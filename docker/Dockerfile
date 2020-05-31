FROM golang:latest

RUN mkdir /go/src/sendgrid

WORKDIR /go/src/sendgrid

RUN go get github.com/sendgrid/sendgrid-go
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers

ADD . /go/src/sendgrid
