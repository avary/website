FROM golang:latest

LABEL maintainer="Bilal Khan"

WORKDIR /website

COPY go.mod go.sum

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main" ]