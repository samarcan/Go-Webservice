package database

import (
	"fmt"
	"os"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Load Postgres integration
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	usermodels "WebService/apps/users/ormModels"
)

var db *gorm.DB

func init() {
	initDatabase()
	db.Debug().AutoMigrate(
		&usermodels.UserORM{},
	)
}

func initDatabase() {
	dir, err := os.Getwd()
	fmt.Println(dir)
	e := godotenv.Load("settings.env")
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")


	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

}

// GetDB Returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}