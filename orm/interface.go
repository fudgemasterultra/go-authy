package orm

import (
	"fmt"
	"fudgemasterultra/go-authy/cli"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func createConnectionString(dbEvn cli.SetDBEnvPath) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbEvn.Host, dbEvn.User, dbEvn.Password, dbEvn.DBName, dbEvn.Port)
}

func IntialSetup(dbEvn cli.SetDBEnvPath){
	connectionString := createConnectionString(dbEvn)
	_, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	os.Setenv("GO_AUTHY_DB_URL", connectionString)

}