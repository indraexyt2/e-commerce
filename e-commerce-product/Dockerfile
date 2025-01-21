# Stage 1: Builder
FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .
COPY .env .

RUN go build -o e-commerce-product
RUN chmod +x e-commerce-product

# Stage 2: Runtime
FROM alpine:3.18

COPY --from=builder /app/e-commerce-product .
COPY --from=builder /app/.env .

EXPOSE 9001

CMD [ "./e-commerce-product" ]