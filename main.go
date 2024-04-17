package main

import (
	"fmt"
	"superapp/app/board"
	"superapp/app/todo"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func setupLogger() *zap.Logger {
	// Configure the logger
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.StacktraceKey = ""

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	return logger
}

func main() {
	engine := html.New("./views", ".html")

	// Initialize Zap logger
	logger := setupLogger()
	defer logger.Sync()

	app := fiber.New(fiber.Config{

		Views: engine,
	})

	// Middleware
	// Middleware: Use Zap logger
	app.Use(
		func(c *fiber.Ctx) error {
			// Log request details
			logger.Info("Request",
				zap.String("method", c.Method()),
				zap.String("path", c.Path()),
				zap.String("ip", c.IP()),
			)
			return c.Next()
		})
	app.Use(recover.New())

	app.Static("/static", "./static")

	// Routes
	todo.Setup(app)
	board.Setup(app)

	// Run the server on port 8000
	err := app.Listen(":8000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
