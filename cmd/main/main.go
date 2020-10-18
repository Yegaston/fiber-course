package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yegaston/fiber-course"
)

func main() {
	app := fiber.New()
	
	app.Group("/movies", webservices.SearchMovieHandler)

	_ = app.Listen(":3000")
}