FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-croud-live .

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /go-croud-live /app/go-croud-live

EXPOSE 8000

CMD ["/app/go-croud-live"]
