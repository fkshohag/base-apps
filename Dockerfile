FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

# COPY go.mod go.sum ./
RUN go mod tidy



RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/*.go

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]