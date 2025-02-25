FROM golang:1.22-alpine as builder

WORKDIR /app
COPY . .

RUN apk add --no-cache \
    sqlite-dev \
    build-base

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /app/finance ./cmd/server

FROM golang:1.22-alpine
WORKDIR /root/

RUN apk add --no-cache sqlite-libs

COPY --from=builder /app/finance ./finance
COPY --from=builder /app/finance.db ./finance.db

RUN chmod +x ./finance

CMD ["./finance"]