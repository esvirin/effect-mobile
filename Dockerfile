FROM golang:1.23-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o subscription-service ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/subscription-service .
CMD ["./subscription-service"]
