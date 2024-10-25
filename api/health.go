package api

import (
	"github.com/gofiber/fiber/v2"
	"go-my-life/internal/domain/entity/health"
	"go-my-life/internal/domain/repository/healthdb"
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

	app.Put("/weight", func(ctx *fiber.Ctx) error {
		weight := &health.Weight{}
		err := ctx.BodyParser(weight)
		if err != nil {
			return err
		}
		weight.UserId = ctx.Locals("userId").(int64)

		err = weight.Put()
		if err != nil {
			return err
		}
		return ctx.JSON(&fiber.Map{
			"success": true,
		})
	})

	app.Post("/weight/query", func(ctx *fiber.Ctx) error {
		param := &healthdb.WeightQueryParam{}
		err := ctx.BodyParser(param)
		if err != nil {
			return err
		}

		param.UserId = ctx.Locals("userId").(int64)
		weights, total := healthdb.QueryWeights(*param)

		return ctx.JSON(&fiber.Map{
			"success": true,
			"data":    weights,
			"total":   total,
		})
	})

	return app
}
