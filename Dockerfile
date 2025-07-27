# Étape 1 : build l'exécutable
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o stargate ./cmd/stargate

# Étape 2 : image légère pour exécuter l'app
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/stargate .
COPY config.yaml .

EXPOSE 8080

CMD ["./mrp"]
