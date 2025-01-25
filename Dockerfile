#Build
FROM golang:1.20-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o receipt-inator ./main.go

#Run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/receipt-inator .
ENV GIN_MODE=release
EXPOSE 8080
CMD ["./receipt-inator"]