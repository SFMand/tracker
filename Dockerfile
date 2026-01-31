FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY main.go .

RUN go mod init tracker
RUN go build -o tracker main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/tracker .

EXPOSE 8080

CMD ["./tracker"]
