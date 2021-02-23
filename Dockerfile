FROM golang:1.15.7-buster

WORKDIR /app

RUN go get github.com/cespare/reflex

CMD ["make", "run"]
