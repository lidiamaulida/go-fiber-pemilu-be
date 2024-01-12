package migrate

import (
	"pemiluApi-go/database"
	"pemiluApi-go/models"
	"fmt"
)

func RunMigrations() {
	err := database.DB.AutoMigrate(&models.Article{},&models.User{},&models.Paslon{},&models.Partai{},&models.Voter{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration success broour")
}