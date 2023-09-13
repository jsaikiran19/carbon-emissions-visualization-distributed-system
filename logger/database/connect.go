package database

import (
    "os"
    "fmt"
    "log"   
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB;
func Connect() {
    
	dsn := get_dsn();
	fmt.Println(dsn);
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});

	if err!=nil {
		panic("Connection to DB failed");
	}
	DB = connection;
    fmt.Println("Connected to PostgreSQL!")
}

func get_dsn() string {

	err := godotenv.Load()
	if err != nil {
		log.Panic("Failed to load ENV")
	}
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")

	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DBNAME := os.Getenv("POSTGRES_DBNAME")

	return fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s port=%s  sslmode=require",
		POSTGRES_HOST, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DBNAME, POSTGRES_PORT)
}
