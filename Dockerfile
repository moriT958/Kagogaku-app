FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .


FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

RUN mkdir -p /app/images

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]