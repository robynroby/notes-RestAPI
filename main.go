package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/robynroby/go-Fiber/database"
	"github.com/robynroby/go-Fiber/notes"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/notes", note.GetNotes)
	app.Get("/note/:id", note.GetNote)
	app.Post("/note", note.NewNote)
	app.Delete("/note/:id", note.DeleteNote)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "note.db")
	if err != nil {
		panic("Failed to connect to Database")
	}
	fmt.PrintLn("Database successfully opened. ")

	database.DBConn.AutoMigrate(&note.Note{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}
