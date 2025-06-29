# EventFlow

[![Go Version](https://img.shields.io/badge/Go-1.24.3+-00ADD8?style=flat&logo=go)](https://golang.org)
[![CI Status](https://img.shields.io/github/actions/workflow/status/TheFoxKD/eventflow/ci.yml?branch=main)](https://github.com/TheFoxKD/eventflow/actions)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/TheFoxKD/eventflow)](https://goreportcard.com/report/github.com/TheFoxKD/eventflow)

EventFlow is a microservice system for event management and real-time notifications. Users can create events, subscribe to categories, and receive notifications through multiple channels.

## 🚀 Tech Stack

- **Backend**: Go 1.24, Clean Architecture
- **HTTP**: Chi Router, JWT Authentication
- **Database**: PostgreSQL with migrations (golang-migrate)
- **Messaging**: Kafka (events), RabbitMQ (notifications)
- **Communication**: gRPC with Protocol Buffers
- **Containers**: Docker, Docker Compose
- **Orchestration**: Kubernetes (KIND), Helm Charts
- **Observability**: Prometheus, Grafana, Loki

## 🏗️ Architecture

```mermaid
graph TB
    Client[Client App] --> API[HTTP API Gateway]
    API --> Auth[Auth Service]
    API --> Events[Events Service]
    API --> Sub[Subscriptions Service]

    Events --> Kafka[Kafka Events]
    Kafka --> Consumer[Event Consumer]
    Consumer --> RabbitMQ[RabbitMQ Notifications]
    RabbitMQ --> Email[Email Service]
    RabbitMQ --> SMS[SMS Service]

    Auth --> DB[(PostgreSQL)]
    Events --> DB
    Sub --> DB

    API --> Prometheus[Prometheus Metrics]
    Consumer --> Prometheus
    Prometheus --> Grafana[Grafana Dashboard]
```

### Clean Architecture Layers

```mermaid
graph LR
    subgraph "External"
        HTTP[HTTP Handlers]
        gRPC[gRPC Handlers]
        Kafka[Kafka Consumer]
    end

    subgraph "Use Cases"
        AuthUC[Auth Usecase]
        EventUC[Event Usecase]
        SubUC[Subscription Usecase]
    end

    subgraph "Repository"
        UserRepo[User Repository]
        EventRepo[Event Repository]
        SubRepo[Subscription Repository]
    end

    subgraph "Entity"
        User[User]
        Event[Event]
        Subscription[Subscription]
    end

    HTTP --> AuthUC
    gRPC --> EventUC
    Kafka --> SubUC

    AuthUC --> UserRepo
    EventUC --> EventRepo
    SubUC --> SubRepo

    UserRepo --> User
    EventRepo --> Event
    SubRepo --> Subscription
```

## 🐳 Quick Start with Docker

```bash
# Clone repository
git clone https://github.com/TheFoxKD/eventflow.git
cd eventflow

# Copy environment file
cp .env.example .env

# Start all services
make docker-up

# Run database migrations
make migrate-up

# Check health
curl http://127.0.0.1:8080/health
```

## 📡 API Endpoints

### Authentication

```bash
# Register
curl -X POST http://127.0.0.1:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'

# Login
curl -X POST http://127.0.0.1:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'

# Refresh token
curl -X POST http://127.0.0.1:8080/auth/refresh \
  -H "Content-Type: application/json" \
  -d '{"refresh_token":"your_refresh_token"}'
```

### Events

```bash
# Create event
curl -X POST http://127.0.0.1:8080/events \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title":"Tech Meetup","description":"Go microservices discussion","category":"tech"}'

# Get events
curl -H "Authorization: Bearer $JWT_TOKEN" \
  http://127.0.0.1:8080/events?page=1&limit=10

# Get event by ID
curl -H "Authorization: Bearer $JWT_TOKEN" \
  http://127.0.0.1:8080/events/{event_id}
```

### Subscriptions

```bash
# Subscribe to category
curl -X POST http://127.0.0.1:8080/subscriptions \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"category":"tech"}'

# Get user subscriptions
curl -H "Authorization: Bearer $JWT_TOKEN" \
  http://127.0.0.1:8080/subscriptions

# Unsubscribe
curl -X DELETE http://127.0.0.1:8080/subscriptions/tech \
  -H "Authorization: Bearer $JWT_TOKEN"
```

## 🛠️ Development Setup

### Prerequisites

- Go 1.24
- Docker & Docker Compose
- Make

### Local Development

```bash
# Install dependencies
go mod download

# Run tests
make test

# Run linter
make lint

# Build application
make build

# Run locally (requires external services)
make run
```

### Database Migrations

```bash
# Create new migration
make migrate-create name=add_events_table

# Apply migrations
make migrate-up

# Rollback migration
make migrate-down
```

## 🔧 Configuration

Environment variables (see `.env.example`):

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | PostgreSQL host | `127.0.0.1` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `DB_NAME` | Database name | `eventflow` |
| `JWT_SECRET` | JWT signing secret | `your-secret-key` |
| `KAFKA_BROKERS` | Kafka broker addresses | `127.0.0.1:9092` |
| `RABBITMQ_URL` | RabbitMQ connection URL | `amqp://127.0.0.1:5672` |

## 🧪 Testing

```bash
# Run all tests
make test

# Run with coverage
make test-coverage

# Run integration tests
make test-integration

# Run specific package tests
go test ./internal/usecase/event -v
```

## 🚢 Deployment

### Docker

```bash
# Build image
docker build -t eventflow:latest .

# Run container
docker run -p 8080:8080 --env-file .env eventflow:latest
```

### Kubernetes

```bash
# Deploy to KIND cluster
make k8s-deploy

# Port forward for testing
kubectl port-forward svc/eventflow 8080:80
```

## 📊 Monitoring

- **Metrics**: <http://127.0.0.1:3000> (Grafana)
- **Logs**: Grafana Loki integration
- **Health**: <http://127.0.0.1:8080/health>
- **Prometheus**: <http://127.0.0.1:9090>
