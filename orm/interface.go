package orm

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func saltShaker() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic("System's secure random is not working")
	}
	return base64.URLEncoding.EncodeToString(b)
}

func hashPassword(password string) string {
	salt := saltShaker()
	passSalted := password + salt
	hasher := sha256.New()
	hasher.Write([]byte(passSalted))
	return string(hasher.Sum(nil))

}

func migrations(db *gorm.DB) {
	db.AutoMigrate(&User{}, &SessionToken{})
}

func createConnectionString(host string, user string, password string, dbName string, port string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
}

func IntialSetup(host string, user string, password string, dbName string, port string) {
	connectionString := createConnectionString(host, user, password, dbName, port)
	fmt.Println(connectionString)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	os.Setenv("GO_AUTHY_DB_URL", connectionString)

	migrations(db)

}

func CreateUser(username string, password string) {
	conString, setupFinished := os.LookupEnv("GO_AUTHY_DB_URL")
	fmt.Println(conString)
	if !setupFinished {
		panic("Database not setup")
	}
	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	hashedPassword := hashPassword(password)
	db.Create(&User{Username: username, Password: hashedPassword})
}
