FROM golang:1.20

WORKDIR /app

RUN go mod init

COPY . .

RUN go build -o fi-ms

CMD ["./fi-ms"]