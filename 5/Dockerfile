FROM golang:1.15

WORKDIR /app

COPY . .

RUN go build -o http-server ./cmd/server/main.go

EXPOSE 8080

CMD ["./cmd/server/main.go"]
