package board

import (
	"github.com/gofiber/fiber/v2"
	//
)

var boards []Board = []Board{
	Board{
		ID:          1,
		Name:        "Test Board",
		Link:        "test",
		Description: "This is a test board",
		Threads: []Thread{
			Thread{
				ID:   1,
				Link: "test",
				Posts: []Post{
					Post{
						ID:      1,
						Link:    "test",
						Content: "This is a test post",
					},
				},
			},
		},
	},
}


func Setup(app *fiber.App) {
	board := app.Group("/board")

	board.Get("/", func(c *fiber.Ctx) error {
		return c.Render("board/index", fiber.Map{
			"boards": boards,
		})
	})

	board.Get("/:link", func(c *fiber.Ctx) error {
		link := c.Params("link")
		for _, board := range boards {
			if board.Link == link {
				return c.Render("board/board", fiber.Map{
					"board": board,
				})
			}
		}
		return c.SendStatus(404)
	})

	board.Get("/thread/:id", func(c *fiber.Ctx) error {
		return c.Render("board/thread", fiber.Map{})
	})

	board.Get("/thread/:id/post", func(c *fiber.Ctx) error {
		return c.Render("board/post", fiber.Map{})
	})

	board.Get("/thread/:id/post/:id", func(c *fiber.Ctx) error {
		return c.Render("board/post", fiber.Map{})
	})

}
