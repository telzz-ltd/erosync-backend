FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/app /app/cmd/api

FROM debian:bookworm-slim

COPY --from=builder /app/app /app

EXPOSE 8080

ENTRYPOINT ["/app"]