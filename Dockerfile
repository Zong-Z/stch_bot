FROM golang:1.15-alpine3.12

ADD . /go/src/app
WORKDIR /go/src/app

RUN go get
RUN go install

RUN go build -o main .

CMD /go/src/app/main
