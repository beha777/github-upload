package main

import (
	"testNEW/db"
	"testNEW/routes"
	"testNEW/settings"
)

func main() {
	settings.AppSettings = settings.ReadSettings()
	db.ConnectDatabase()

	/*var users []models.User
	db.GetDBConn().Find(&users)
	for _, val := range users {
		err := bcrypt.CompareHashAndPassword([]byte(val.Password), []byte("qwerty"))
		if err != nil {
			panic(err)
		} else {
			fmt.Println(true)
		}
	}*/
	routes.Init()


}
