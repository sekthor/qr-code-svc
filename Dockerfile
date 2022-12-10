FROM golang:1.17-alpine

WORKDIR /app

COPY ./ ./
RUN go mod download
RUN go build -o /qrcode-service ./cmd/qrcode/main.go
EXPOSE 8080
CMD ["/qrcode-service"]
