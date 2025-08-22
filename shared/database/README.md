# PostgreSQL Database Module

This shared module provides PostgreSQL database connectivity for the standard-fiber-web-api project using GORM.

## Features

- PostgreSQL connection management with GORM
- Environment-based configuration
- Connection pooling
- Database health checks
- Auto-migration support
- Transaction support

## Configuration

Set the following environment variables (see `env_example` for reference):

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password_here
DB_NAME=standard_fiber_db
DB_SSL_MODE=disable
DB_TIMEZONE=UTC
DB_MAX_OPEN_CONNS=25
DB_MAX_IDLE_CONNS=5
```

## Usage

### Initialize Database Connection

```go
import "standard-fiber-go/shared/database"

func main() {
    // Initialize database connection
    err := database.InitializeDatabase()
    if err != nil {
        log.Fatal("Failed to initialize database:", err)
    }
    defer database.CloseDatabase()
    
    // Your application code here
}
```

### Get Database Instance

```go
import "standard-fiber-go/shared/database"

func someFunction() {
    db := database.GetDB()
    // Use db for your queries
}
```

### Auto Migration

```go
import (
    "standard-fiber-go/shared/database"
    userModels "standard-fiber-go/shared/models/user"
)

func migrateModels() {
    err := database.AutoMigrate(&userModels.User{})
    if err != nil {
        log.Fatal("Migration failed:", err)
    }
}
```

### Using Transactions

```go
import "standard-fiber-go/shared/database"

func someTransactionalOperation() error {
    return database.Transaction(func(tx *gorm.DB) error {
        // Your transactional code here
        return nil
    })
}
```

### Health Checks

The module provides database health check endpoints:

- `GET /api/v1/health/database` - Returns detailed database health information

## Example Model

```go
package userModels

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID        uint           `json:"id" gorm:"primarykey"`
    Username  string         `json:"username" gorm:"uniqueIndex;not null"`
    Email     string         `json:"email" gorm:"uniqueIndex;not null"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
```

## Dependencies

- `gorm.io/gorm` - GORM ORM library
- `gorm.io/driver/postgres` - PostgreSQL driver for GORM

## Structure

```
shared/
├── config/
│   └── database.go          # Database configuration
├── database/
│   ├── connection.go        # Connection management
│   └── postgres.go          # PostgreSQL utilities
└── models/
    └── user/
        └── user.go          # Example user model
```