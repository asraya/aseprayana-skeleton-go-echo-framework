# GO IMAGE
FROM golang:1.21.0-alpine3.16

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .
RUN go build -o seeder .

CMD ["./seeder"]

RUN go mod download