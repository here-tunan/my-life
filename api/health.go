package api

import (
	"github.com/gofiber/fiber/v2"
	"go-my-life/internal/service"
)

func HealthMount() *fiber.App {
	app := fiber.New()

	app.Post("/physical/list", func(ctx *fiber.Ctx) error {
		param := &service.PhysicalQueryParam{}
		err := ctx.BodyParser(param)
		if err != nil {
			return err
		}
		res, total := service.QueryPhysicals(param)
		return ctx.JSON(&fiber.Map{
			"success": true,
			"data":    res,
			"total:":  total,
		})
	})

	return app
}
