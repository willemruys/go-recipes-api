package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var db *gorm.DB

func SetupModels() *gorm.DB {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var host string
	var user string
	var password string
	var databaseName string 

	socketDir := os.Getenv("DB_SOCKET_DIR")

	if (len(socketDir) == 0) {
		socketDir = "/cloudsql"
	}

	if os.Getenv("ENVIRONMENT") == "LOCAL" {
		host = os.Getenv("POSTGRES_HOST_LOCAL")
		user = os.Getenv("POSTGRES_USER_LOCAL")
		password = os.Getenv("POSTGRES_PASSWORD_LOCAL")
		databaseName =  os.Getenv("POSTGRES_DB_LOCAL")
		// port = os.Getenv("POSTGRES_PORT_LOCAL")
	}

	if os.Getenv("ENVIRONMENT") == "LOCAL_GCL" {
		host = os.Getenv("POSTGRES_HOST_LOCAL_GCL")
		user = os.Getenv("POSTGRES_USER_LOCAL_GCL")
		password = os.Getenv("POSTGRES_PASSWORD_LOCAL_GCL")
		databaseName =  os.Getenv("POSTGRES_DB_LOCAL_GCL")
		// port = os.Getenv("POSTGRES_PORT_LOCAL_GCL")
	}

	if os.Getenv("ENVIRONMENT") == "DEV" {
		host = socketDir + "/" + os.Getenv("POSTGRES_HOST_DEV")
		user = os.Getenv("POSTGRES_USER_DEV")
		password = os.Getenv("POSTGRES_PASSWORD_DEV")
		databaseName =  os.Getenv("POSTGRES_DB_DEV")
		// port = os.Getenv("POSTGRES_PORT_DEV")
	}
	
	if os.Getenv("ENVIRONMENT") == "PROD" {
		host = os.Getenv("POSTGRES_HOST_PROD")
		user = os.Getenv("POSTGRES_USER_PROD")
		password = os.Getenv("POSTGRES_PASSWORD_PROD")
		databaseName =  os.Getenv("POSTGRES_DB_PROD")
		// port = os.Getenv("POSTGRES_PORT_PROD")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, databaseName)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connecting to DB")
	}

	db.AutoMigrate(&User{}, &Recipe{}, &Comment{}, &List{})

	return db

}

func LoadDB() *gorm.DB {
	return db
}

func SetupMockModels() *gorm.DB {
	var err error

	err = godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	var host string
	var user string
	var password string
	var databaseName string 
	var port string

	if os.Getenv("ENVIRONMENT") == "LOCAL" {
		host = os.Getenv("POSTGRES_HOST_LOCAL")
		user = os.Getenv("POSTGRES_USER_LOCAL")
		password = os.Getenv("POSTGRES_PASSWORD_LOCAL")
		databaseName =  os.Getenv("POSTGRES_DB_LOCAL")
		port = os.Getenv("POSTGRES_PORT_LOCAL")
	}

	if os.Getenv("ENVIRONMENT") == "LOCAL_GCL" {
		host = os.Getenv("POSTGRES_HOST_LOCAL")
		user = os.Getenv("POSTGRES_USER_LOCAL")
		password = os.Getenv("POSTGRES_PASSWORD_LOCAL")
		databaseName =  os.Getenv("POSTGRES_DB_LOCAL")
		port = os.Getenv("POSTGRES_PORT_LOCAL")
	}

	if os.Getenv("ENVIRONMENT") == "DEV" {
		host = os.Getenv("POSTGRES_HOST_DEV")
		user = os.Getenv("POSTGRES_USER_DEV")
		password = os.Getenv("POSTGRES_PASSWORD_DEV")
		databaseName =  os.Getenv("POSTGRES_DB_DEV")
		port = os.Getenv("POSTGRES_PORT_DEV")
	}
	
	if os.Getenv("ENVIRONMENT") == "PROD" {
		host = os.Getenv("POSTGRES_HOST_PROD")
		user = os.Getenv("POSTGRES_USER_PROD")
		password = os.Getenv("POSTGRES_PASSWORD_PROD")
		databaseName =  os.Getenv("POSTGRES_DB_PROD")
		port = os.Getenv("POSTGRES_PORT_PROD")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, databaseName, port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connecting to DB")
	}

	db.Migrator().DropTable(&User{}, &Recipe{}, &Comment{}, &List{})

	db.AutoMigrate(&User{}, &Recipe{}, &Comment{}, &List{})


	return db
}