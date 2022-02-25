package database

import (
	"errors"
	"memorise/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Connect() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:@tcp(localhost:3306)/memorise"), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to DB")
	}
}

func AutoMigrate(){
	DB.AutoMigrate(models.User{})
	DB.AutoMigrate(models.Feedback{})
	DB.AutoMigrate(models.Gallery{})
	DB.AutoMigrate(models.Admin{})
	DB.Migrator().HasTable(models.Admin{});{
		if err:=DB.First(&models.Admin{}).Error; errors.Is(err, gorm.ErrRecordNotFound){
			admin := models.Admin{
				Email: "admin@gmail.com",
			}
			admin.SetPassword("administrator")
			DB.Create(&admin)
		}
	}
}