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
	app := fiber.New()

	// Default config
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	initDatabase()
	// defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(8000)
}

// import (
// 	"fmt"

// 	"github.com/gofiber/fiber"
// 	"github.com/glebarez/sqlite"
// 	"github.com/robynroby/go-Fiber/database"
// 	"github.com/robynroby/go-Fiber/notes"
// 	"gorm.io/gorm"
// )

// func setupRoutes(app *fiber.App) {
// 	app.Get("/notes", note.GetNotes)
// 	app.Get("/note/:id", note.GetNote)
// 	app.Post("/note", note.NewNote)
// 	app.Delete("/note/:id", note.DeleteNote)
// }

// func initDatabase() {
// 	var err error
// 	database.DBConn, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("Failed to connect to Database")
// 	}
// 	fmt.PrintLn("Database successfully opened. ")

// 	database.DBConn.AutoMigrate(&note.Note{})
// 	fmt.Println("Database Migrated")
// }

// func main() {
// 	app := fiber.New()

// 	initDatabase()
// 	defer database.DBConn.Close()

// 	setupRoutes(app)

// 	app.Listen(3000)
// }
