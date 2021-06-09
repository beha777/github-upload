package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testNEW/models"
	"testNEW/settings"
)

var database *gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
	Serial uint
}

func initDB() *gorm.DB {
	settingParams := settings.AppSettings.PostgresParams
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		settingParams.Server, settingParams.Port,
		settingParams.User, settingParams.DataBase,
		settingParams.Password)
	db, err := gorm.Open("postgres", connString)

	if err != nil {
		log.Fatal("Couldn't connect to postgresql database", err.Error())
		//logger.Error.Println()
	}
	db.LogMode(true)
	db.SingularTable(true)

	db.AutoMigrate(&models.User{})


	return db
}

func ConnectDatabase () {
	database = initDB()
}

func GetDBConn() *gorm.DB {
	return database
}

func IsExistUser(login string) bool {
	var user models.User
	var count int
	GetDBConn().First(&user, models.User{Login: login}).Count(&count)
	return count > 0 && user.ID > 0
}

func HashAllPasswords() {
	var users []models.User
	for _, val := range users {
		bytes, err := bcrypt.GenerateFromPassword([]byte(val.Password), 14)
		if err != nil {
			log.Fatal("Cant hash password", err)
		}
		database.Table("user").Where("login = ?", val.Login).Update("password", string(bytes))
	}
}