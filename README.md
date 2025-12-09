# Payment Gateway
A microservice built in Go for processing credit card transactions with a robust architecture following clean code principles.
## Architecture
The project follows a clean architecture pattern with the following layers:
- **Domain**: Business entities and rules
- **Repository**: Data persistence layer
- **Use Cases**: Business logic
- **Infrastructure**: External dependencies (database, APIs)
## Tech Stack
- Go 1.15
- PostgreSQL
- Docker & Docker Compose
## Prerequisites
- Docker
- Docker Compose
- Git
## Getting Started
### 1. Clone the repository
```bash
git clone git@github.com:brandaoplaster/payment-gateway.git
cd payment-gateway
```
### 2. Configure environment variables
Create a `.env` file in the root directory:
```env
# Database
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=payment_gateway
```
### 3. Start the application with Docker Compose
```bash
docker compose up -d
```
This will start:
- PostgreSQL database on port 5432
- Payment Gateway API on port 8080
- PgAdmin on port 9000
  - Access: http://localhost:9000
  - Email: admin@user.com
  - Password: 123456
### 4. Check if services are running
```bash
docker compose ps
```
### 5. View logs
```bash
# All services
docker compose logs -f
# Specific service
docker compose logs -f payment_gateway
```
## Development
### Running without Docker
```bash
# Install dependencies
go mod download
# Run the application
go run cmd/server/main.go
```
### Run tests
```bash
go test ./...
```
### Format code
```bash
go fmt ./...
```
### Build binary
```bash
go build -o payment-gateway cmd/server/main.go
```
## Docker Commands
### Stop services
```bash
docker compose down
```
### Stop and remove volumes
```bash
docker compose down -v
```
### Rebuild images
```bash
docker compose build --no-cache
docker compose up -d
```
### Access database
```bash
docker compose exec postgres psql -U postgres -d payment_gateway
```
## Project Structure
```
payment-gateway/
├── cmd/              # Application entry points
│   └── server/
│       └── main.go
├── domain/           # Business entities and rules
│   ├── credit_card.go
│   └── transaction.go
├── repository/       # Data access layer
│   └── transaction_repository.go
├── usecase/          # Business logic
│   └── process_transaction.go
├── infrastructure/   # External dependencies
│   └── database/
│       └── postgres.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```
## License
This project is licensed under the MIT License - see the LICENSE file for details.
