package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ProvideDB initializes the GORM database connection for PostgreSQL.
func ProvideDB() (*gorm.DB, error) {
	// Load environment variables or use default values
	host := os.Getenv("POSTGRESQL_DB_HOST")
	// if host == "" {
	// 	host = "localhost"
	// }
	port := os.Getenv("POSTGRESQL_DB_PORT")
	// if port == "" {
	// 	port = "5432"
	// }
	user := os.Getenv("POSTGRESQL_DB_USERNAME")
	// if user == "" {
	// 	user = "postgres"
	// }
	password := os.Getenv("POSTGRESQL_DB_PASSWORD")
	// if password == "" {
	// 	password = "Ak.123456Qw"
	// }
	dbname := os.Getenv("POSTGRESQL_DB_NAME")
	// if dbname == "" {
	// 	dbname = "crossword"
	// }

	// PostgreSQL connection string
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		host, port, user, password, dbname,
	)

	// Configure GORM with PostgreSQL driver
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
		return nil, err
	}

	// Test the database connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
