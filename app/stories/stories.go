package stories

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var tasks = []string{"Task one", "Task two", "Task three"}

func Setup(app *fiber.App) {
	todo := app.Group("/stories")
	todo.Get("/", func(c *fiber.Ctx) error {
		return c.Render("todo/index", fiber.Map{})
	})

	todo.Get("/tasks", func(c *fiber.Ctx) error {
		return c.Render("todo/list", fiber.Map{"Tasks": tasks})
	})

	todo.Post("/add", func(c *fiber.Ctx) error {
		task := c.FormValue("task")
		tasks = append(tasks, task)
		fmt.Println(tasks)
		return c.Redirect("/todo/tasks")
	})

}
