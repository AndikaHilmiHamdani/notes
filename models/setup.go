package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// var dbUsername = os.Getenv("DB_USERNAME")
	// var dbPassword = os.Getenv("DB_PASSWORD")
	// var dbHost = os.Getenv("DB_HOST")
	// var dbPort = os.Getenv("DB_PORT")
	// var dbName = os.Getenv("DB_NAME")

	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/users_db"))

	// db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	DB = db

}
