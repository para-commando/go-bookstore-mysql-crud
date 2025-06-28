# Managing Secrets in Go Applications

This document outlines the best practices for managing secrets and configuration in Go applications.

## 1. Environment Variables (Current Implementation)

### Setup
1. Create a `.env` file in your project root (never commit this file):
```bash
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_actual_password
DB_NAME=BOOK_STORE

# Application Configuration
APP_PORT=8080
APP_ENV=development
```

2. Add `.env` to your `.gitignore`:
```gitignore
.env
*.env
```

3. Use the configuration in your application:
```go
config := config.LoadConfig()
dsn := config.GetDSN()
```

### Running the Application
```bash
# Set environment variables directly
export DB_PASSWORD=your_password
go run cmd/main/main.go

# Or use a .env file with a library like godotenv
```

## 2. Alternative Approaches

### A. Using godotenv Library
```bash
go get github.com/joho/godotenv
```

```go
package main

import (
    "github.com/joho/godotenv"
    "log"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using system environment variables")
    }
}
```

### B. Configuration Files (for non-sensitive data)
```go
type Config struct {
    Database DatabaseConfig `json:"database"`
    Server   ServerConfig   `json:"server"`
}
```

### C. Secret Management Services

#### AWS Secrets Manager
```go
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/secretsmanager"
)

func getSecret(secretName string) (string, error) {
    svc := secretsmanager.New(session.Must(session.NewSession()))
    input := &secretsmanager.GetSecretValueInput{
        SecretId: aws.String(secretName),
    }
    
    result, err := svc.GetSecretValue(input)
    if err != nil {
        return "", err
    }
    
    return *result.SecretString, nil
}
```

#### HashiCorp Vault
```go
import (
    vault "github.com/hashicorp/vault/api"
)

func getVaultSecret(path string) (map[string]interface{}, error) {
    client, err := vault.NewClient(vault.DefaultConfig())
    if err != nil {
        return nil, err
    }
    
    secret, err := client.Logical().Read(path)
    if err != nil {
        return nil, err
    }
    
    return secret.Data, nil
}
```

### D. Kubernetes Secrets
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: db-secret
type: Opaque
data:
  username: cm9vdA==  # base64 encoded
  password: cGFzc3dvcmQ=  # base64 encoded
```

```go
// In your Go application
dbUser := os.Getenv("DB_USER")
dbPassword := os.Getenv("DB_PASSWORD")
```

## 3. Security Best Practices

### ✅ DO:
- Use environment variables for secrets
- Never commit secrets to version control
- Use different secrets for different environments
- Rotate secrets regularly
- Use strong, unique passwords
- Encrypt secrets at rest
- Use secret management services in production

### ❌ DON'T:
- Hardcode secrets in source code
- Commit `.env` files to version control
- Use the same secrets across environments
- Log secrets or sensitive data
- Store secrets in client-side code
- Use weak passwords

## 4. Production Deployment

### Docker
```dockerfile
FROM golang:1.24-alpine
WORKDIR /app
COPY . .
RUN go build -o main cmd/main/main.go

# Don't copy .env files
CMD ["./main"]
```

```bash
# Run with environment variables
docker run -e DB_PASSWORD=secret -e DB_HOST=prod-db your-app
```

### Kubernetes
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: bookstore-app
spec:
  template:
    spec:
      containers:
      - name: app
        image: your-app:latest
        env:
        - name: DB_PASSWORD
          valueFrom:
            secretKeyRef:
              name: db-secret
              key: password
```

## 5. Development Setup

1. Copy `config.example` to `.env`
2. Update the values in `.env` with your actual credentials
3. Never commit the `.env` file
4. Use different credentials for development, staging, and production

## 6. Testing

For testing, you can use environment variables or mock configurations:

```go
func TestDatabaseConnection(t *testing.T) {
    // Set test environment variables
    os.Setenv("DB_HOST", "localhost")
    os.Setenv("DB_PASSWORD", "test_password")
    
    config := LoadConfig()
    // ... test logic
}
``` 