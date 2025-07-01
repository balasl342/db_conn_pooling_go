# Go DB Connection Pooling Example

This is a simple Go program that demonstrates how to configure and force database connection pooling using the `database/sql` standard package with PostgreSQL.

## 🧪 What It Does

- Sets up a connection pool with 50 max open connections
- Launches 50 concurrent goroutines
- Each goroutine runs a query to force connection usage
- Shows how `database/sql` lazily opens connections

## 🔧 Configuration

Update your DB credentials inside `main.go`:

```go
const (
    dbHost     = "localhost"
    dbPort     = 5432
    dbUser     = "postgres"
    dbPassword = "yourpassword"
    dbName     = "yourdb"
)
