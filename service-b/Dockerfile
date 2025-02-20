FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o service-b ./cmd/server


FROM debian:bookworm-slim

WORKDIR /root/

COPY --from=builder /app/service-b .

ENV SERVICE_A_URL=http://service-a:8080
ENV RABBITMQ_URI=amqp://guest:guest@rabbitmq:5672/

EXPOSE 50051

CMD ["./service-b"]
