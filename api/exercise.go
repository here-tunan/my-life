package api

import (
	"github.com/gofiber/fiber/v2"
	"go-my-life/internal/domain/entity/exercise"
	service "go-my-life/internal/service/exercise"
)

func ExerciseMount() *fiber.App {
	app := fiber.New()

	app.Put("", func(ctx *fiber.Ctx) error {
		theExercise := &exercise.Log{}
		err := ctx.BodyParser(theExercise)
		if err != nil {
			return err
		}
		//res := service.PutExercise(ctx.Locals("userId").(int64), theExercise)
		res := service.PutExercise(2, theExercise)
		return ctx.JSON(&fiber.Map{
			"success": true,
			"data":    res,
		})
	})

	return app
}
