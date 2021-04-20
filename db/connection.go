package db
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"user_api/model"
)

var DB *gorm.DB

func ConnectDatabase () {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_tutorial?charset=utf8&parseTime=True")
	if err != nil {
		panic("Fail to connect DB")
	}

	db.AutoMigrate(&model.User{},&model.Employee{})
	DB = db
}