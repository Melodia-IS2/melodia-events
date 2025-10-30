FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o melodia-events ./cmd/service

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app/

COPY --from=builder /app/melodia-events .
COPY --from=builder /app/docs/kafka/kafka-init.sh ./scripts/kafka-init.sh

EXPOSE 8085

CMD ["./melodia-events"]

