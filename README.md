# TOTM API (Theatre of the Mind API)

A Go-based REST API for the Theatre of the Mind RPG platform, providing authentication, blog management, and various RPG tools.

## Overview

TOTM API is a structured, production-ready Go application built with modern practices including:
- **Repository Pattern** for data access
- **Dependency Injection** for service management
- **Structured Logging** with Zap
- **JWT Authentication**
- **Database Migrations** with Flyway
- **Comprehensive Testing** (Unit & Integration)

## Architecture

The application follows a clean architecture pattern with clear separation of concerns:

```
cmd/main.go                 # Application entry point
├── internal/               # Private application code
│   ├── config/            # Configuration management
│   ├── controllers/       # HTTP request handlers
│   ├── db/               # Data access layer (repositories)
│   ├── di/               # Dependency injection container
│   ├── dto/              # Data Transfer Objects
│   ├── logger/           # Structured logging
│   ├── middleware/       # HTTP middleware
│   ├── models/           # Database models
│   └── nithian/          # Custom RPG tools
├── migrations/           # Database schema migrations
└── protos/              # Protocol buffer definitions
```

## Packages Documentation

### Core Application

#### `cmd/`
**Purpose**: Application entry point and server initialization
- **`main.go`**: Bootstraps the application, initializes logging, database connection, and starts the HTTP server
- Handles environment configuration and graceful shutdown

#### `internal/config/`
**Purpose**: Configuration management and database connection settings
- **`DbConfig.go`**: Database configuration structure and connection string management
- **`DbConfig_test.go`**: Unit tests for configuration functionality

### Data Layer

#### `internal/models/`
**Purpose**: Database entity models representing the database schema
- **`Users.go`**: User entity with authentication fields (username, password, salt, role, etc.)
- **`Blogposts.go`**: Blog post entity with content, metadata, and publishing status
- **`Person.go`**: Person entity for testing and reference
- **`Shops.go`**: Shop and inventory management entities
- **`ShopStock.go`**: Shop inventory tracking
- **`FlywaySchemaHistory.go`**: Migration tracking table

#### `internal/db/`
**Purpose**: Data access layer implementing the Repository pattern
- **`IDbContext.go`**: Interface defining database operations
- **`DbContext.go`**: Concrete implementation using SQLx and Squirrel query builder
- **`UserRepository.go`**: User data operations (CRUD, authentication queries)
- **`BlogpostRepository.go`**: Blog post data operations
- **`mocks/`**: Mock implementations for testing
  - **`MockDbContext.go`**: Mock database context for unit tests
  - **`MockDbContext_test.go`**: Tests for mock functionality

#### `internal/dto/`
**Purpose**: Data Transfer Objects for API communication
- **`User.go`**: User DTO for API requests/responses
- **`BlogPost.go`**: Blog post DTO
- **`converter.go`**: Generic conversion utilities between models and DTOs
- **`converter_test.go`**: Unit tests for conversion logic

### Business Logic

#### `internal/controllers/`
**Purpose**: HTTP request handlers and business logic
- **`registry.go`**: Route registration system for modular controllers

**Auth Controller (`auth/`)**:
- **`AuthService.go`**: JWT-based authentication, user registration, password management
- **`dto.go`**: Authentication-specific DTOs (login, registration, password change)

**Blog Controller (`blog/`)**:
- **`blog_service.go`**: Blog post CRUD operations, content management

**Health Controller (`health/`)**:
- **`HealthCheck.go`**: Application health monitoring, database connectivity checks

**OpenAI Controller (`open_ai/`)**:
- **`openai.go`**: Integration with OpenAI API for AI-powered features

**Shop Controller (`shop/`)**:
- **`ShopController.go`**: Shop and inventory management
- **`dto.go`**: Shop-related DTOs
- **`testdata.go`**: Test data for shop functionality

### Infrastructure

#### `internal/di/`
**Purpose**: Dependency injection container for service management
- **`container.go`**: Service registration and resolution, database connection management
- **`container_test.go`**: Integration tests for DI container

#### `internal/logger/`
**Purpose**: Structured logging with Zap
- **`logger.go`**: Logger initialization, configuration, and helper functions
- **`logger_test.go`**: Unit tests for logging functionality
- Features: JSON output, function tracing, contextual fields, error handling

#### `internal/middleware/`
**Purpose**: HTTP middleware for cross-cutting concerns
- **`jwtAuth.go`**: JWT token validation and user authentication
- **`cors.go`**: Cross-Origin Resource Sharing configuration
- **`sanitiser.go`**: Input sanitization and XSS protection
- **`basicAuth.go`**: Basic authentication middleware

### Specialized Features

#### `internal/nithian/`
**Purpose**: Custom RPG tools and utilities
- **`translator.go`**: Nithian language translation system
- **`translator_test.go`**: Unit tests for translation functionality

### Database Management

#### `migrations/`
**Purpose**: Database schema versioning and migrations
- **`totmapi/flyway.toml`**: Flyway configuration
- **`totmapi/migrations/`**: SQL migration files (V001-V011)
- **`totmapi/schema-model/`**: Generated schema models

### API Definitions

#### `protos/`
**Purpose**: Protocol buffer definitions for future gRPC support
- **`chat.proto`**: Chat service definitions for real-time messaging

## Key Features

### Authentication & Security
- JWT-based authentication with configurable expiration
- Password hashing with SHA-256 and salt
- Role-based access control
- Input sanitization and XSS protection
- CORS configuration

### Data Management
- Repository pattern for clean data access
- Generic DTO conversion system
- SQL query building with Squirrel
- Database migrations with Flyway
- Comprehensive error handling

### Observability
- Structured JSON logging with Zap
- Function-level tracing
- Contextual logging fields
- Health check endpoints
- Database connectivity monitoring

### Testing
- Unit tests with build tags (`//go:build unit`)
- Integration tests with build tags (`//go:build integration`)
- Mock implementations for isolated testing
- Database integration tests

## Technology Stack

- **Language**: Go 1.23.2
- **Web Framework**: Gorilla Mux
- **Database**: PostgreSQL with SQLx
- **Query Builder**: Squirrel
- **Authentication**: JWT (golang-jwt)
- **Logging**: Zap (Uber)
- **Testing**: Testify
- **Migrations**: Flyway
- **AI Integration**: OpenAI API
- **Input Sanitization**: Bluemonday

## Getting Started

### Prerequisites
- Go 1.23.2+
- PostgreSQL
- Environment variable: `TOTM_CONN_STRING`

### Running the Application
```bash
# Set database connection string
export TOTM_CONN_STRING="postgres://user:pass@localhost:5432/dbname?sslmode=disable"

# Run the application
go run ./cmd/main.go

# Or build and run
go build ./cmd/main.go
./main.exe
```

### Testing
```bash
# Run unit tests
go test ./... -tags=unit

# Run integration tests
go test ./... -tags=integration

# Run all tests
make all
```

### Database Migrations
The application uses Flyway for database migrations. Migrations are automatically applied on startup.

## API Endpoints

### Authentication
- `POST /login` - User login
- `POST /register` - User registration
- `POST /changepass` - Password change

### Blog Management
- `GET /blogposts` - List published blog posts
- `GET /blogposts/{id}` - Get specific blog post
- `POST /blogposts` - Create new blog post
- `PUT /blogposts` - Update blog post

### Health & Monitoring
- `GET /health` - Application health check

### AI Features
- `POST /openai/prompt` - OpenAI integration

### Shop Management
- `GET /shop` - Shop information

## Development

### Project Structure
The codebase follows Go project layout conventions with clear separation between public and internal packages.

### Code Quality
- Comprehensive test coverage
- Structured logging throughout
- Error handling with proper context
- Clean architecture principles
- Dependency injection for testability

### Contributing
1. Follow existing code patterns
2. Add tests for new functionality
3. Use structured logging
4. Update documentation as needed
5. Run all tests before committing

## License

[Add your license information here] 