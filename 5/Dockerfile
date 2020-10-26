FROM golang:1.15

ENV GO111MODULE=on

  ARG GOFLAGS=-mod=vendor
  ENV GOFLAGS="${GOFLAGS}"

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o 5 http-server/cmd/server/main.go    

CMD ["./http-server"]
