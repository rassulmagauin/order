FROM golang:1.22 AS builder

WORKDIR /app

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o order ./cmd/order/

FROM alpine:latest  

COPY --from=builder /app/order /order

CMD ["/order"]
