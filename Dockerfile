FROM golang:1.18 as build

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o cutter ./cmd/app/main.go

ENTRYPOINT ["/app/cutter"]

EXPOSE 8080
