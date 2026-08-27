FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app 

FROM alpine:3.14

COPY --from=builder /app/app /app

EXPOSE 8080

ENTRYPOINT ["/app"]