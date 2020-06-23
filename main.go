package main

import (
	"fmt"

	"github.com/fabiendeborde/noaproject_api/database"
	"github.com/fabiendeborde/noaproject_api/project"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


func setupRoutes(app *fiber.App) {
	app.Get("api/v1/project", project.GetProjects)
	app.Get("api/v1/project/:id", project.GetProject)
	app.Post("api/v1/project", project.NewProject)
	app.Delete("api/v1/project/:id", project.DeleteProject)
	app.Post("api/v1/project/:id/like", project.AddLikeProject)
	app.Post("api/v1/project/:id/unlike", project.RemoveLikeProject)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "projects.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Databse connection successfully opened")

	database.DBConn.AutoMigrate(&project.Project{})
	fmt.Println("Databse migrated")

}

func main() {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) {
    // IsWebSocketUpgrade returns true if the client
    // requested upgrade to the WebSocket protocol.
    if websocket.IsWebSocketUpgrade(c) {
      c.Locals("allowed", true)
      c.Next()
    }
  })
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}