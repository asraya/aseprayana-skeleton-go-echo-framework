package configs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"aseprayana-skeleton-go/entity"
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create the connection string
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open connection to the database
	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// SetConnIdleTime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(10 * time.Hour)

	// Migrate the schema
	db.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.RoleUser{},
		&entity.Limit{},
		&entity.Transaction{},
		&entity.Goods{},
	)

	return db
}

func InitialMigration(db *gorm.DB) {
	// Get the underlying sql.DB instance from the gorm.DB instance.
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	dbSQL.Close() // Close the database connection
}
