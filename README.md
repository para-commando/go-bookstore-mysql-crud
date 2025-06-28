# Go Bookstore API

A RESTful API for managing books in a bookstore, built with Go, GORM, and MySQL.

## 🚀 Features

- **CRUD Operations**: Create, Read, Update, Delete books
- **RESTful API**: Follows REST principles
- **Database Integration**: MySQL with GORM ORM
- **Swagger Documentation**: Interactive API documentation
- **Environment Configuration**: Secure secret management
- **Docker Support**: Containerized deployment

## 📋 Prerequisites

- Go 1.24 or higher
- MySQL 8.0 or higher
- Make (optional, for using Makefile commands)

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-bookstore
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Setup environment**
   ```bash
   # Copy environment template
   cp config.example .env
   
   # Edit .env with your database credentials
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=your_password
   DB_NAME=BOOK_STORE
   APP_PORT=8080
   APP_ENV=development
   ```

4. **Install development tools (optional)**
   ```bash
   make install-tools
   ```

## 🏃‍♂️ Running the Application

### Using Makefile (Recommended)
```bash
# Run the application
make run

# Run in development mode
make run-dev

# Run in production mode
make run-prod
```

### Using Go directly
```bash
go run cmd/main/main.go
```

## 📚 API Documentation

### Swagger UI
Once the application is running, you can access the interactive API documentation at:
```
http://localhost:8080/swagger/
```

### Available Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/books` | Get all books |
| GET | `/books/{id}` | Get a book by ID |
| POST | `/books` | Create a new book |
| PUT | `/books/{id}` | Update a book |
| DELETE | `/books/{id}` | Delete a book |

### Example API Usage

#### Create a Book
```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "price": "$15.99"
  }'
```

#### Get All Books
```bash
curl http://localhost:8080/books
```

#### Get Book by ID
```bash
curl http://localhost:8080/books/1
```

#### Update a Book
```bash
curl -X PUT http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Great Gatsby (Updated)",
    "author": "F. Scott Fitzgerald",
    "price": "$19.99"
  }'
```

#### Delete a Book
```bash
curl -X DELETE http://localhost:8080/books/1
```

## 🔧 Development

### Generate Swagger Documentation
```bash
# Generate Swagger docs
make swagger-init

# Or manually
swag init -g cmd/main/main.go -o ./docs
```

### Run Tests
```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run short tests
make test-short
```

### Code Quality
```bash
# Format code
make fmt

# Vet code
make vet

# Run linter
make lint

# Run all checks
make check
```

### Build
```bash
# Build for current platform
make build

# Build for specific platforms
make build-linux
make build-windows
make build-mac

# Build release binaries
make release
```

## 🐳 Docker

### Build Docker Image
```bash
make docker-build
```

### Run Docker Container
```bash
make docker-run
```

### Clean Docker Images
```bash
make docker-clean
```

## 📁 Project Structure

```
go-bookstore/
├── cmd/
│   └── main/
│       └── main.go          # Application entry point
├── pkg/
│   ├── config/
│   │   ├── app.go          # Database connection
│   │   └── config.go       # Configuration management
│   ├── controllers/
│   │   └── bookstore-controller.go  # HTTP handlers
│   ├── models/
│   │   └── book.go         # Data models
│   ├── routes/
│   │   └── bookstore-router.go      # Route definitions
│   └── utils/
│       └── bookstore-utils.go       # Utility functions
├── docs/                   # Generated Swagger documentation
├── .env                    # Environment variables (not in git)
├── config.example          # Environment template
├── Makefile               # Build and development commands
├── go.mod                 # Go module file
└── README.md              # This file
```

## 🔒 Security

- Environment variables for sensitive data
- No hardcoded credentials
- Proper error handling
- Input validation

See [SECRETS.md](SECRETS.md) for detailed security guidelines.

## 🧪 Testing

The application includes comprehensive tests for all endpoints and business logic.

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./pkg/controllers -v
```

## 📊 API Response Format

### Success Response
```json
{
  "id": 1,
  "title": "The Great Gatsby",
  "author": "F. Scott Fitzgerald",
  "price": "$15.99",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### Error Response
```json
{
  "error": "Book not found"
}
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Run the test suite
6. Submit a pull request

## 📄 License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

For support and questions:
- Create an issue in the repository
- Check the [Swagger documentation](http://localhost:8080/swagger/) when running
- Review the [SECRETS.md](SECRETS.md) for configuration help

## 🔄 Version History

- **v1.0.0**: Initial release with CRUD operations and Swagger documentation 