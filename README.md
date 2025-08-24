# CMS CISDI API

## 🚀 Features

- **RESTful API** with comprehensive endpoints
- **Database Migration** with Goose integration
- **Redis Caching** for improved performance
- **JWT Authentication** for secure access
- **Docker Support** for easy deployment
- **API Documentation** with Swagger/OpenAPI
- **Health Check** endpoints for monitoring
- **Structured Logging** for better observability
- **Hot Reload** support for development
- **Comprehensive Testing** with coverage reports

## 🛠️ Tech Stack

- **Backend**: Go 1.23
- **Database**: PostgreSQL 15
- **Cache**: Redis 7
- **Documentation**: Swagger/OpenAPI
- **Migration**: Goose
- **Testing**: Testify
- **Containerization**: Docker & Docker Compose
- **Hot Reload**: Air

## 📋 Prerequisites

Before running this project, make sure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.23 or later)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/) (optional but recommended)

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/dimasbayuseno/cisdi-go-test.git
cd cisdi-go-test
```

### 2. Environment Setup

Create a `.env` file in the root directory:

```bash
cp .env.example .env
```

Configure your environment variables:

```env
# Application Configuration
APP_PORT=8080
APP_ENV=development

# Database Configuration
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=cms_cisdi_db
DB_SSL_MODE=disable

# Redis Configuration
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key
JWT_EXPIRATION=72h
```

### 3. Run with Docker Compose (Recommended)

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f app

# Stop services
docker-compose down
```

### 4. Run Locally (Development)

```bash
# Install dependencies
go mod download

# Run database and Redis with Docker
docker-compose up -d db redis

# Update .env for local development
# Change DB_HOST=localhost and REDIS_HOST=localhost

# Run database migrations
make migrate-up

# Generate API documentation
make swag

# Start the application
make run
```

## 📚 API Documentation

Once the application is running, you can access:

- **API Documentation**: http://localhost:8080/swagger/index.html
- **Health Check**: http://localhost:8080/api/health-check

## 🔧 Development

### Available Make Commands

```bash
# Code Generation
make mocks          # Generate mocks for testing
make swag          # Generate Swagger documentation

# Development
make run           # Run the application locally
make build         # Build the application
make build-win     # Build for Windows
make air           # Run with hot reload (Unix)
make air-win       # Run with hot reload (Windows)

# Database
make migrate-up    # Run database migrations
make migrate-down  # Rollback migrations
make migrate-new   # Create fresh migration

# Testing
make test-cov      # Run tests with coverage report

# Docker
make docker-run    # Run Docker container manually
```

### Hot Reload Development

For development with automatic reloading:

```bash
# Install Air (if not already installed)
go install github.com/cosmtrek/air@latest

# Run with hot reload
make air        # For Unix systems
make air-win    # For Windows
```

### Database Migrations

Create a new migration:

```bash
make migrate-new
# Follow the prompts to create migration files
```

Run migrations:

```bash
make migrate-up    # Apply pending migrations
make migrate-down  # Rollback last migration
```

## 🧪 Testing

Run the test suite:

```bash
# Run tests with coverage
make test-cov

# View coverage report
open cover.html
```

The project includes:
- Unit tests for business logic
- Integration tests for database operations
- API endpoint testing
- Mock generation for dependencies

## 🐳 Docker Deployment

### Build and Deploy

```bash
# Build the Docker image
docker-compose build

# Start all services
docker-compose up -d

# Scale the application (if needed)
docker-compose up -d --scale app=3
```

### Environment-specific Deployment

For different environments, create environment-specific compose files:

```bash
# Production
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d

# Staging
docker-compose -f docker-compose.yml -f docker-compose.staging.yml up -d
```

## 📊 Monitoring & Health Checks

### Health Check Endpoint

```bash
curl http://localhost:8080/api/health-check
```

Response:

```"json
  "OK"
```

### Logging

The application uses structured logging with different levels:

- **Error**: Critical errors that need immediate attention
- **Warn**: Warning messages for potential issues
- **Info**: General information about application flow
- **Debug**: Detailed information for debugging

Logs are written to both console and file (`logs/app.log`).

## 👥 Authors

- **Dimas Bayu Suseno** - [GitHub](https://github.com/dimasbayuseno)
