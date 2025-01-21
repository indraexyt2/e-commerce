# Stage 1: Builder
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .
COPY .env .

RUN go build -o e-commerce-order
RUN chmod +x e-commerce-order

# Stage 2: Runtime
FROM alpine:3.18

COPY --from=builder /app/e-commerce-order .
COPY --from=builder /app/.env .

EXPOSE 9002

CMD [ "./e-commerce-order" ]