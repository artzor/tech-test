FROM golang:1.14.2-alpine3.11 AS builder
LABEL stage=builder
WORKDIR /
COPY . .

RUN GOOS=linux go build -o app cmd/main.go

FROM alpine:3.11.6 as final
WORKDIR /
COPY --from=builder /app .
ENTRYPOINT ["/app"]
