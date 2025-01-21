# 🛒 E-Commerce Microservices Project

This project is a microservices-based e-commerce platform built using **GO**, **Echo**, **Gin**, **GORM**, **Docker**, **Redis**, **Kafka**, **PostgreSQL**, and **MySQL**. It consists of multiple services that handle user management, product management, order processing, and payment processing. The project also integrates with an external e-wallet service for payment transactions.

---

## 📋 Table of Contents

1. [Project Overview](#-project-overview)
2. [Services](#-services)
3. [Prerequisites](#-prerequisites)
4. [Setup and Run](#-setup-and-Run)
5. [Environment Variables](#-environment-variables)
7. [API Endpoints](#-api-endpoints)

---

## 🚀 Project Overview

This project is designed to simulate a real-world e-commerce platform using microservices architecture. Each service is containerized using Docker and communicates with other services via REST APIs, Redis, and Kafka. The services are:

- **User Management Service (UMS)**: Handles user authentication, registration, and profile management.
- **Product Service**: Manages product listings, inventory, and product details.
- **Order Service**: Handles order creation, status updates, and order history.
- **Payment Service**: Processes payments and integrates with an external e-wallet service.
- **E-Wallet Service**: An external service that handles wallet transactions (linked to the payment service).

---

## 🛠 Services

### 1. **User Management Service (UMS)**
- **Port**: `9000`
- **Database**: PostgreSQL (`e_commerce_ums`)
- **Dependencies**: PostgreSQL

### 2. **Product Management Service**
- **Port**: `9001`
- **Database**: PostgreSQL (`e_commerce_product`)
- **Dependencies**: PostgreSQL, Redis (for caching)

### 3. **Order Management Service**
- **Port**: `9002`
- **Database**: PostgreSQL (`e_commerce_order`)
- **Dependencies**: PostgreSQL, Kafka (for event-driven communication)

### 4. **Payment Management Service**
- **Port**: `9003`
- **Database**: PostgreSQL (`e_commerce_payment`)
- **Dependencies**: PostgreSQL, Kafka, E-Wallet Service

### 5. **E-Wallet Service**
- **Port**: `8081`
- **Database**: MySQL (`e_wallet_wallet`)
- **Dependencies**: MySQL

---

## 📦 Prerequisites

Before running the project, ensure you have the following installed:

- **Docker**: [Install Docker](https://docs.docker.com/get-docker/)
- **Docker Compose**: [Install Docker Compose](https://docs.docker.com/compose/install/)
- **Git**: [Install Git](https://git-scm.com/downloads)

---

## 🛠 Setup and Run

1. Clone the **e-commerce** repository:
   ```bash
   git clone https://github.com/indraexyt2/e-commerce.git
   cd e-commerce

2. Navigate to the e-commerce-infra to run the project with Docker Compose:
    ```bash
   cd e-commerce-infra
   docker-compose up -d

3. Clone the **e-wallet** repository (required for the payment service):
    ```bash
   git clone https://github.com/indraexyt2/e-wallet.git
   cd e-wallet-infra

4. Navigate to the e-commerce-infra to run the project with Docker Compose:
    ```bash
   cd e-wallet-infra
   docker-compose up -d

## 🔧 Environment Variables

Each service requires specific environment variables to function correctly. Below are the required variables for each service:

### User Management Service (UMS)
```bash
    PORT=9000
    DB_HOST=postgres
    DB_PORT=5432
    DB_NAME=e_commerce_ums
    DB_USER=postgres
    DB_PASSWORD=postgres
    JWT_SECRET=secret
   ```
### Product Service

```bash
   PORT=9001
   DB_HOST=postgres
   DB_PORT=5432
   DB_NAME=e_commerce_product
   DB_USER=postgres
   DB_PASSWORD=postgres
   REDIS_HOST=redis-1:6371
   UMS_URL=http://e-commerce-ums:9000
   UMS_ENDPOINT_PROFILE=/user/v1/profile
```

### Order Service

```bash
    PORT=9002
    DB_HOST=postgres
    DB_PORT=5432
    DB_NAME=e_commerce_order
    DB_USER=postgres
    DB_PASSWORD=postgres
    KAFKA_HOST=kafka1:9092,kafka2:9093,kafka3:9094
    KAFKA_TOPIC_PAYMENT_INITIATE=payment-initiate-topic
    KAFKA_TOPIC_PAYMENT_REFUND=payment-refund-topic
    UMS_URL=http://e-commerce-ums:9000
    UMS_ENDPOINT_PROFILE=/user/v1/profile
```

### Payment Service

```bash
    PORT=9003
    DB_HOST=postgres
    DB_PORT=5432
    DB_NAME=e_commerce_payment
    DB_USER=postgres
    DB_PASSWORD=postgres
    KAFKA_HOST=kafka1:9092,kafka2:9093,kafka3:9094
    KAFKA_TOPIC_PAYMENT_INITIATE=payment-initiate-topic
    KAFKA_TOPIC_PAYMENT_INITIATE_PARTITION=3
    KAFKA_TOPIC_PAYMENT_REFUND=payment-refund-topic
    KAFKA_TOPIC_PAYMENT_REFUND_PARTITION=3
    UMS_URL=http://e-commerce-ums:9000
    UMS_ENDPOINT_PROFILE=/user/v1/profile
    WALLET_URL=http://wallet:8081
    WALLET_ENDPOINT_PAYMENT_LINK=/wallet/v1/ex/link
    WALLET_ENDPOINT_PAYMENT_UNLINK=/wallet/v1/ex/%d/unlink
    WALLET_ENDPOINT_PAYMENT_LINK_CONFIRM=/wallet/v1/ex/link/%d/confirmation
    WALLET_ENDPOINT_PAYMENT_TRANSACTION=/wallet/v1/ex/transaction
    WALLET_SECRET_KEY=secret-key
    WALLET_CLIENT_ID=e-commerce
    E_COMMERCE_URL=http://e-commerce-order:9002
    E_COMMERCE_ENDPOINT_ORDER_CALLBACK=/order/v1/in/%d/status
```

### E-Wallet Service

```bash
    PORT=8081
    DB_HOST=mysql
    DB_PORT=3306
    DB_NAME=e_wallet_wallet
    DB_USER=root
    DB_PASSWORD=root
```

## 📡 API Endpoints

### 👤 User Management Service (UMS) API Endpoints

| **Method** | **Endpoint**               | **Middleware**               |
|------------|----------------------------|-------------------------------|
| `POST`     | `/user/v1/register`        | -                             |
| `POST`     | `/user/v1/register/admin`  | -                             |
| `POST`     | `/user/v1/login`           | -                             |
| `POST`     | `/user/v1/login/admin`     | -                             |
| `GET`      | `/user/v1/profile`         | `MiddlewareValidateAuth`      |
| `PUT`      | `/user/v1/refresh-token`   | `MiddlewareRefreshToken`      |
| `DELETE`   | `/user/v1/logout`          | `MiddlewareValidateAuth`      |

### 📦 Product Service API Endpoints

#### Product Endpoints
| **Method** | **Endpoint**                     | **Middleware**               |
|------------|----------------------------------|-------------------------------|
| `POST`     | `/product/v1`                   | `MiddlewareValidateAuth`      |
| `PUT`      | `/product/v1/:id`               | `MiddlewareValidateAuth`      |
| `PUT`      | `/product/v1/variant/:id`       | `MiddlewareValidateAuth`      |
| `DELETE`   | `/product/v1/:id`               | `MiddlewareValidateAuth`      |
| `GET`      | `/product/v1/list`              | -                             |
| `GET`      | `/product/v1/:id`               | -                             |

#### Category Endpoints
| **Method** | **Endpoint**                     | **Middleware**               |
|------------|----------------------------------|-------------------------------|
| `POST`     | `/product/v1/category`          | `MiddlewareValidateAuth`      |
| `PUT`      | `/product/v1/category/:id`      | `MiddlewareValidateAuth`      |
| `DELETE`   | `/product/v1/category/:id`      | `MiddlewareValidateAuth`      |
| `GET`      | `/product/v1/category`          | -                             |

### 📦 Order Service API Endpoints

| **Method** | **Endpoint**                     | **Middleware**               |
|------------|----------------------------------|-------------------------------|
| `POST`     | `/order/v1`                     | `MiddlewareValidateAuth`      |
| `PUT`      | `/order/v1/:id/status`          | `MiddlewareValidateAuth`      |
| `PUT`      | `/order/v1/in/:id/status`       | -                             |
| `GET`      | `/order/v1/:id`                 | `MiddlewareValidateAuth`      |
| `GET`      | `/order/v1`                     | `MiddlewareValidateAuth`      |

### 💳 Payment Service API Endpoints

| **Method** | **Endpoint**                     | **Middleware**               |
|------------|----------------------------------|-------------------------------|
| `POST`     | `/payment/v1/link`              | `MiddlewareValidateAuth`      |
| `POST`     | `/payment/v1/link/confirm`      | `MiddlewareValidateAuth`      |
| `POST`     | `/payment/v1/unlink`            | `MiddlewareValidateAuth`      |