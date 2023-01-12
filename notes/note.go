package note

import (
	"github.com/gofiber/fiber"
	"github.com/robynroby/go-Fiber/database"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

func GetNotes(c *fiber.Ctx) {
	db := database.DBConn
	var notes []Note
	db.Find(&notes)
	c.JSON(notes)
}

func GetNote(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var note Note
	db.Find(&note, id)
	c.JSON(note)
}

func NewNote(c *fiber.Ctx) {
	db := database.DBConn

	note := new(Note)
	if err := c.BodyParser(note); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&note)
	c.JSON(note)
}

func DeleteNote(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var note Note
	db.First(&note, id)

	if note.Title == "" {
		c.Status(500).Send("No note found with the given ID")
		return
	}
	db.Delete(&note)
	c.Send("Book successfully deleted")
}
