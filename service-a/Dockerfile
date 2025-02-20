# Imagen base con Go
FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o service-a .

# Imagen final con Debian más reciente
FROM debian:bookworm-slim

WORKDIR /root/

COPY --from=builder /app/service-a .

COPY --from=builder /app/config/config.yml ./config/config.yml

ENV MONGO_URI=mongodb://mongo:27017
ENV RABBITMQ_URI=amqp://guest:guest@rabbitmq:5672/

EXPOSE 8080

CMD ["./service-a"]
