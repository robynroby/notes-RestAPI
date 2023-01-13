package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/robynroby/go-Fiber/database"
	note "github.com/robynroby/go-Fiber/notes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// fmt.PrintLn("Database successfully opened. ")

	database.DBConn.AutoMigrate(&note.Note{})
	fmt.Println("Database Migrated")
}

func main() {

	initDatabase()
	// defer database.DBConn.Close()

	app := fiber.New()

	app.Use(cors.New()) // CORS middleware

	setupRoutes(app)

	app.Listen(8000)
}
