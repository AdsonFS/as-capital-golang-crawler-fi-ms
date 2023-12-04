FROM golang:1.20

WORKDIR /app

RUN go mod init crwler-fi

COPY . .

RUN go build -o fi-ms cmd/main.go 

CMD ["./fi-ms"]