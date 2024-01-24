FROM golang:1.21

WORKDIR /app

RUN go mod init crawler-fi

COPY . .

RUN go build -o fi-ms cmd/main.go

ENV REDIS_ADDR=redis:6379
ENV REDIS_PASS=

CMD ./fi-ms -redis=$REDIS_ADDR -pass=$REDIS_PASS
