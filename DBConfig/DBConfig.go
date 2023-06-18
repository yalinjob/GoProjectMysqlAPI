package dbConfig

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbUser string
var doPass string
var dbHost string
var dbName string
var dbPort string
var dataSourceName string
var dbValidConnection bool

var (
	db *gorm.DB
)

// Create a MySQL database connection string

//dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

// Open a connection to the MySQL database

func DbConnectionVerification() {

	//dbUser := "root"
	//dbPass := "UziNarkis5!"
	//dbName := "GoUserList"
	//dbHost := "localhost"
	//dbPort := 3306

	//dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	d, err := gorm.Open("mysql", "root:UziNarkis5!@tcp(localhost:3306)/GoUserList?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}

	db = d
	fmt.Println("Successfully connected to the MySQL database!")

}
func GetDB() *gorm.DB {
	return db
}
