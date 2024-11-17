package connector

import (
	"SerieForge/entities"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

var tables []interface{} = []interface{}{
	entities.Series{}, entities.Season{}, entities.Episode{},
}

func InitializeDBConnection() bool {
	var err error
	dsn := os.Getenv("SQL_CON_STRING")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: GetCustomDbLogger(),
	})

	if err != nil {
		fmt.Println("Error in connecting to the database")
		fmt.Println(err)
	}

	createEntities()
	return true
}

func createEntities() {
	if err := Db.AutoMigrate(tables...); err != nil {
		fmt.Println("Error in migration")
		fmt.Println(err)
	}
}

func CloseDbConnection() error {
	DB, _ := Db.DB()
	return DB.Close()
}
