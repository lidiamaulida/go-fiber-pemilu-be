package main

import(
	"pemiluApi-go/database"
	"pemiluApi-go/migrate"
	"github.com/gofiber/fiber/v2"
	"pemiluApi-go/controllers/articlescontrollers"
	"pemiluApi-go/controllers/userscontrollers"
	"pemiluApi-go/controllers/paslonscontrollers"
	"pemiluApi-go/controllers/partaiscontrollers"
	"pemiluApi-go/controllers/voterscontrollers"


) 

func main() {
	database.ConnectiDatabase()
	migrate.RunMigrations()


	app := fiber.New()

	//article
	app.Get("/api/articles", articlescontrollers.Getall)
	app.Get("/api/article/:id", articlescontrollers.GetOne)

	//user
	app.Get("/api/users", userscontrollers.Getall)
	app.Get("/api/user/:id", userscontrollers.GetOne)
	app.Post("/api/user", userscontrollers.Create)

	//paslon
	app.Get("/api/paslons", paslonscontrollers.Getall)
	app.Get("/api/paslon/:id", paslonscontrollers.GetOne)
	app.Post("/api/paslon", paslonscontrollers.Create)

	//partai
	app.Get("/api/partai", partaiscontrollers.Getall)
	app.Get("/api/partai/:id", partaiscontrollers.GetOne)
	app.Post("/api/partai", partaiscontrollers.Create)

	//voter
	app.Get("/api/voters", voterscontrollers.Getall)
	app.Get("/api/voter/:id", voterscontrollers.GetOne)
	app.Post("/api/voter", voterscontrollers.Create)


	err := app.Listen(":8000")
    if err != nil {
      panic(err)
    }
}