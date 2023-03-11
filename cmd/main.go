package main

import (
	"log"

	"github.com/AndryMaullana/go-fiber-api/pkg/common/config"
	"github.com/AndryMaullana/go-fiber-api/pkg/common/db"
	"github.com/go-delve/delve/pkg/config"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db.Init(c.DBUrl)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString(c.Port)
	})
	//  book.RegisterRouter(app, db)
	app.Listen(c.Port)

}
