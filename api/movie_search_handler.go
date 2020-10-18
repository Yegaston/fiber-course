package api

import "github.com/gofiber/fiber"


func (w *WebServices)SearchMovieHandler(c *fiber.Ctx) {
	res, err : = w.S.search.Search(MovieFilter{})

	if err != nil {
		err := fiber.NewError(400, "Cannot bring movies")
		c.Next(err)
	}

	c.Send(res)
}