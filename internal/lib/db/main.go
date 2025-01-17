package db

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database credentials.
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "gosimpleauth"
)

func GetDb() (*gorm.DB, error) {
	// Connection string
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&VerificationToken{}, &Session{}, &OauthToken{}, &User{})
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
}

func InitGormDb() {
	// Initialize database
	db, err := GetDb()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Migrate models
	migrate(db)

	fmt.Println("Database initialized and migrated successfully!")
}

func InsertUser(db *gorm.DB) {
	password := "Karthik@150502"
	user := User{
		FullName:      "Karthik J",
		Email:         "karthikrdy150502@gmail.com",
		Password:      &password,
		EmailVerified: false,
		Role:          "Admin",
	}

	// Insert the record
	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	log.Println("User inserted successfully with ID:", user.ID)
}
