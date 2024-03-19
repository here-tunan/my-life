package api

import (
	"github.com/gofiber/fiber/v2"
	"go-my-life/internal/domain/entity/family"
	"go-my-life/internal/service"
	"go-my-life/pkg/myerror"
	"go-my-life/pkg/myresult"
)

func FamilyMount() *fiber.App {
	app := fiber.New()

	app.Post("", func(ctx *fiber.Ctx) error {
		userId := ctx.Locals("userId").(int64)

		param := &family.Family{}

		// 解析JSON RequestBody
		err := ctx.BodyParser(param)

		if err != nil {
			return ctx.JSON(&fiber.Map{
				"success": false,
				"code":    "500",
				"error":   err.Error(),
			})
		}
		// 创建家庭
		err = service.CreateFamily(userId, param)
		if err != nil {
			return ctx.JSON(&fiber.Map{
				"success": false,
				"code":    "500",
				"error":   err.Error(),
			})
		}

		return ctx.JSON(&fiber.Map{
			"success": true,
			"code":    "200",
			"data":    param,
		})
	})

	app.Get("", func(ctx *fiber.Ctx) error {
		userId := ctx.Locals("userId").(int64)
		res, err := service.GetFamily(userId)
		result := &myresult.MyResult[family.Family]{}

		if err != nil {
			if err.Error() == myerror.NoFamily.String() || err.Error() == myerror.NotFamilyMember.String() {
				result.Err(int(myerror.NoFamily), err.Error())
			} else {
				result.Err(int(myerror.InternalError), err.Error())
			}
		} else {
			result.OK(*res)
		}
		return ctx.JSON(result)
	})

	return app
}
