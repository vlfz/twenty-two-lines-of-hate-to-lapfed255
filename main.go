package main

import (
    "log"
	"github.com/gofiber/fiber"
)

type User struct {
	Name string `json:"name"`
	Message string `json:"message"`
	IsTrue bool `json:"isTrue"`
}

func main() {
    app := fiber.New()

    app.Get("/", func (c *fiber.Ctx) error {
        return c.JSON(User { Name: "cipherskyinc", Message: "Hate lapfed255 and Go", IsTrue: true })
    })

    log.Fatal(app.Listen(":51515"))
}