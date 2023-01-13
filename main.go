package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/robynroby/go-Fiber/database"
	note "github.com/robynroby/go-Fiber/notes"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// fmt.PrintLn("Database successfully opened. ")

	database.DBConn.AutoMigrate(&note.Note{})
	fmt.Println("Database Migrated")
}

func main(){
	app := fiber.New()
	setupRoutes(app)

	app.Use(cors.New())
	initDatabase()
	app.Listen(8000)
}




// func main() {

// 	initDatabase()
// 	// defer database.DBConn.Close()

// 	app := fiber.New()

// 	// default fiber cors
// 	app.Use(cors.New())
	
// 	// app.Use(middleware.Logger())
// 	setupRoutes(app)

// 	app.Listen(8000)
// }
