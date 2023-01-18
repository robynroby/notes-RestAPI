package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/robynroby/go-Fiber/database"
	note "github.com/robynroby/go-Fiber/notes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/middleware"
)

func setupRoutes(app *fiber.App) {
	app.Get("/notes", note.GetNotes)
	app.Get("/notes/:id", note.GetNote)
	app.Post("/notes", note.NewNote)
	app.Delete("/notes/:id", note.DeleteNote)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to Database")
	}

	fmt.Println("Database successfully opened. ")

	database.DBConn.AutoMigrate(&note.Note{})
	fmt.Println("Database Migrated")
}

func main() {
	// initilize database
	initDatabase()

	// create new fiber app
	app := fiber.New()

	// default fiber cors
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// setup routes
	setupRoutes(app)

	// listen on port 8000
	app.Listen(8000)

}
