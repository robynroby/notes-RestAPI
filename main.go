package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/robynroby/go-Fiber/database"
	note "github.com/robynroby/go-Fiber/notes"
	"gorm.io/driver/mysql"
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
	database.DBConn, err = gorm.Open(mysql.Open("root:root@notes/notes"), &gorm.Config{})
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
		AllowCredentials: true,
	}))
	// setup routes
	setupRoutes(app)

	// listen on port 8000
	app.Listen(8000)

}
