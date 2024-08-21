package orm

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
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

func hashPassword(password string) (hash []byte, salt string) {
	salt = saltShaker()
	passSalted := password + salt
	hasher := sha256.New()
	hasher.Write([]byte(passSalted))
	fmt.Println(salt)
	return hasher.Sum(nil), salt

}

func migrations(db *gorm.DB) {
	db.AutoMigrate(&User{}, &SessionToken{})
}

func usernameTaken(username string, db *gorm.DB) bool {
	var user *User
	result := db.Where("username = ?", username).First(&user)

	return result.Error == nil
}

func findUser(email string, db *gorm.DB) (*User, error) {
	var user *User
	result := db.Where("email = ?", email).First(&user)
	fmt.Println(result.Error)
	return user, result.Error
}

func dbConnect() *gorm.DB {
	var connData ConnectionData
	file, err := os.Open("config.yml")
	if err != nil {
		panic("Config.yml not present. Please run db-setup")
	}
	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	bs := make([]byte, fileStat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		panic(err)
	}
	err = yaml.Unmarshal(bs, &connData)
	if err != nil {
		panic(err)
	}
	conString := createConnectionString(connData)
	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return db

}

func createConnectionString(connInfo ConnectionData) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", connInfo.Host, connInfo.User, connInfo.Password, connInfo.DBName, connInfo.Port)
}

func writeYML(yml []byte) {
	var file *os.File
	var err error
	file, err = os.Create("config.yml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(yml)
}

func createConConfig(connectionInfo ConnectionData) {
	byteYml, err := yaml.Marshal(connectionInfo)
	if err != nil {
		panic(err)
	}
	writeYML(byteYml)
}

func IntialSetup(host string, user string, password string, dbName string, port string) {
	connInfo := ConnectionData{
		Host:     host,
		User:     user,
		Password: password,
		DBName:   dbName,
		Port:     port,
	}
	connectionString := createConnectionString(connInfo)
	fmt.Println(connectionString)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	os.Setenv("GO_AUTHY_DB_URL", connectionString)
	createConConfig(connInfo)
	migrations(db)

}

func CreateUser(email string, username string, password string) {
	db := dbConnect()
	hashedPassword, salt := hashPassword(password)
	user, err := findUser(email, db)
	fmt.Println(user)
	userNameTaken := usernameTaken(username, db)
	if err == nil {
		panic("Email already in use")
	}
	if userNameTaken {
		panic("Username already in use")
	}
	db.Create(&User{Username: username, Email: email, Password: hashedPassword, Salt: salt})
}
