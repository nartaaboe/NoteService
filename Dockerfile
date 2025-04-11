FROM golang:1.18-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o notes-microservice ./cmd/server

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/notes-microservice .

EXPOSE 8080

CMD ["./notes-microservice"]
