FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app ./cmd/app

FROM scratch AS release

COPY --from=builder /bin/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/app"]