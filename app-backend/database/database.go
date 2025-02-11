package database

import (
	"database/sql"
	"log"
	"sync"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DatabaseConfig 
type DatabaseConfig struct {
	User     string
	Password string
	DbName   string
	Host     string
	Port     string
}

// DB pool
type DB struct {
	pool *sql.DB
}

// singleton instance
var (
	instance *DB
	once     sync.Once
)

// NewDB create pool
func NewDB() *DB {
	once.Do(func() {
		user := os.Getenv("PGUSER")
		password := os.Getenv("PGPASSWORD")
		dbName := os.Getenv("PGDATABASE")
		host := os.Getenv("PGHOST")
		port := os.Getenv("PGPORT")

		connStr := "user=" + user +
			" password=" + password +
			" dbname=" + dbName +
			" host=" + host +
			" port=" + port +
			" sslmode=disable"

		pool, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal("Error connecting to the database: ", err)
		}

		// set pool config
		pool.SetMaxOpenConns(25) 
		pool.SetMaxIdleConns(25)  
		pool.SetConnMaxLifetime(0)

		// test connection
		err = pool.Ping()
		if err != nil {
			log.Fatal("Error pinging the database: ", err)
		}

		instance = &DB{pool: pool}
	})

	return instance
}

// GetPool return instance pool
func GetPool() *sql.DB {
	if instance == nil {
		log.Fatal("Database connection pool not initialized")
	}
	return instance.pool
}
