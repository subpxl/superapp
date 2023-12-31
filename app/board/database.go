package board

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Define a package-level variable to hold the database connection

func ConnectDatabase() error {

	dsn := "host=localhost user=fourchango password=fourchango dbname=fourchango port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("db loaded")
	DB = db

	return nil

}

func SyncDb() {

	fmt.Println("db synced")
	DB.AutoMigrate(Board{})
	DB.AutoMigrate(Thread{})
	DB.AutoMigrate(Post{})
}
