FROM golang:1.12.5

RUN apt-get install git

WORKDIR /go/src/github.com/aut-ce/Web101
COPY ./server/ .
RUN rm -rf /data/*

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["Web101"]