package api

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-my-life/internal/domain/repository/userdb"
	"go-my-life/internal/infrastructure"
	"log"
	"strings"
)

var authWhiteList = []string{"login", "validToken", "refreshToken", "category_analysis", "category_immigrate"}

func Start() {

	app := fiber.New(fiber.Config{
		//EnablePrintRoutes: true,
	})

	app.Use(cors.New())

	// 定义一个中间件来进行token验证
	validateMiddleware := func(c *fiber.Ctx) error {
		api := c.Path()
		if isNotNeedAuth(api) {
			return c.Next()
		}

		token := c.Get("Authorization")
		if token == "" {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   "用户未登录",
				"code":    401,
			})
		}

		// 查询token，如果未找到返回一个错误响应
		result, err := infrastructure.Redis.Get(context.Background(), token).Result()
		if err != nil && result == "" {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   "登录token过期，请重新登录",
				"code":    401,
			})
		}

		var userResource userdb.User
		err = json.Unmarshal([]byte(result), &userResource)
		if err != nil {
			println("user unmarshal error: " + err.Error())
		}

		c.Locals("userId", userResource.Id)

		// 如果验证通过，继续处理请求
		return c.Next()
	}
	app.Use(validateMiddleware)

	app.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.SendString("ha!")
	})

	// 设置根路径
	root := app.Group("/api")

	root.Mount("/money", MoneyMount())
	root.Mount("/health", HealthMount())
	root.Mount("/user", UserMount())
	root.Mount("/family", FamilyMount())
	root.Mount("/oss", OssMount())

	log.Fatal(app.Listen(":3001"))
}

func isNotNeedAuth(api string) bool {
	for _, s := range authWhiteList {
		if strings.HasSuffix(api, s) {
			return true
		}
	}
	return false
}
