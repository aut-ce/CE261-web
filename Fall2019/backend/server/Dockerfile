FROM golang:1.12.5-alpine

RUN apk add git

WORKDIR /go/src/web101
COPY ./server/ .
RUN rm -rf /data/*

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["web101"]