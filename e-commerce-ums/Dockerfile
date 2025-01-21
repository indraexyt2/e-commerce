# Stage 1: Builder
FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . .

# Build the application
RUN go build -o e-commerce-ums
RUN chmod +x e-commerce-ums

# Stage 2: Runtime
FROM alpine:3.18

WORKDIR /app

# Copy binary and environment file
COPY --from=builder /app/e-commerce-ums .
COPY --from=builder /app/.env .

EXPOSE 9000

CMD ["./e-commerce-ums"]
