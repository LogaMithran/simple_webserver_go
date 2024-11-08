package connectors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"power-sync/entities"
	"strconv"
)

var DB *gorm.DB

func CreateDataBaseConnection() {
	var err error
	dsn := os.Getenv("SQL_CON_STRING")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: GetCustomGormLogger(),
	})
	if err != nil {
		fmt.Println("Error in connecting to the database")
		fmt.Println(err)
	}
	createEntities()
}

func createEntities() {
	fmt.Println("Checking for entities")

	if err := DB.AutoMigrate(&entities.People{}, &entities.Movie{}); err != nil {
		fmt.Println("Error in migrating the database")
		fmt.Println(err)
	}
}

func Close() error {
	db, _ := DB.DB()
	return db.Close()
}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		limit, _ := strconv.Atoi(c.Query("limit"))
		if limit == 0 {
			limit = 10
		}
		if page == 0 {
			page = 1
		}
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}

}
