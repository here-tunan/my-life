package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-my-life/internal/domain/entity/user"
	"go-my-life/internal/domain/repository/userdb"
	"go-my-life/internal/infrastructure"
	"go-my-life/internal/service"
	"go-my-life/pkg/myerror"
	"go-my-life/pkg/myresult"
	"strconv"
	"time"
)

// 一个小时
const expiredSeconds = 60 * 60

func UserMount() *fiber.App {

	app := fiber.New()

	app.Get("/info", func(ctx *fiber.Ctx) error {
		userId := ctx.Locals("userId").(int64)

		result, err := service.GetUser(&userdb.User{
			Id: userId,
		})
		if err != nil {
			return ctx.JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		return ctx.JSON(&fiber.Map{
			"success": true,
			"data":    result,
		})
	})

	app.Get("/login", func(ctx *fiber.Ctx) error {
		account := ctx.Query("account")
		password := ctx.Query("password")
		fmt.Printf("user login: %s.", account)

		if account == "" || password == "" {
			return ctx.JSON(&fiber.Map{
				"success": false,
				"error":   "登录失败，账号密码错误",
			})
		}

		userParam := &userdb.User{
			Account:  account,
			Password: password,
		}
		userResult, _ := service.GetUser(userParam)

		if userResult != nil {
			userJSON, err := json.Marshal(userResult)
			if err != nil {
				// 处理错误
				return ctx.JSON(&fiber.Map{
					"success": false,
					"error":   err.Error(),
				})
			}

			token := userResult.Account + uuid.New().String()

			cmdStatus := infrastructure.Redis.Set(context.Background(), token, userJSON, time.Duration(expiredSeconds)*time.Second)
			if cmdStatus.Err() != nil {
				return ctx.JSON(&fiber.Map{
					"success": false,
					"error":   cmdStatus.Err().Error(),
				})
			}

			return ctx.JSON(&fiber.Map{
				"success": true,
				"data": map[string]string{
					"token":       token,
					"expiredTime": strconv.Itoa(expiredSeconds),
				},
			})
		}

		return ctx.JSON(&fiber.Map{
			"success": false,
			"error":   "登录失败，账号密码错误",
		})
	})

	app.Get("/validToken", func(ctx *fiber.Ctx) error {
		token := ctx.Query("token")
		result := infrastructure.Redis.Get(context.Background(), token)

		code := 200
		if result.Err() != nil {
			errorMsg := ""
			if result.Val() == "" {
				code = 401
			} else {
				code = 500
				errorMsg = result.Err().Error()
			}
			return ctx.JSON(&fiber.Map{
				"success": false,
				"code":    code,
				"error":   errorMsg,
			})
		}

		if result.Val() != "" {
			return ctx.JSON(&fiber.Map{
				"success": true,
				"code":    code,
			})
		}

		return ctx.JSON(&fiber.Map{
			"success": false,
			"error":   "token 已过期",
			"code":    401,
		})
	})

	app.Get("/refreshToken", func(ctx *fiber.Ctx) error {
		token := ctx.Query("token")
		result := infrastructure.Redis.Get(context.Background(), token)

		if result != nil && result.Val() != "" {
			userJson := result.Val()
			var userResource userdb.User
			err := json.Unmarshal([]byte(userJson), &userResource)
			if err != nil {
				return ctx.JSON(&fiber.Map{
					"success": false,
					"error":   err.Error(),
				})
			}
			newToken := userResource.Account + uuid.New().String()

			infrastructure.Redis.Set(context.Background(), newToken, userJson, time.Duration(expiredSeconds)*time.Second)
			return ctx.JSON(&fiber.Map{
				"success": true,
				"data": map[string]string{
					"token":       newToken,
					"expiredTime": strconv.Itoa(expiredSeconds),
				},
			})
		}

		return ctx.JSON(&fiber.Map{
			"success": false,
			"error":   "token 已过期, 请重新登录",
			"code":    401,
		})
	})

	app.Put("", func(ctx *fiber.Ctx) error {
		userId := ctx.Locals("userId").(int64)
		userInfo := &user.User{}
		result := &myresult.MyResult[user.User]{}
		err := ctx.BodyParser(userInfo)
		userInfo.Id = userId
		if err != nil {
			return ctx.JSON(result.Err(int(myerror.WrongParam), err.Error()))
		}
		// 保存user信息
		err = service.PutUser(userInfo)
		if err != nil {
			return ctx.JSON(result.Err(int(myerror.InternalError), err.Error()))
		}
		return ctx.JSON(result.OK(*userInfo))
	})

	return app
}
