package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-my-life/internal/domain/entity/money"
	"go-my-life/internal/domain/repository/moneydb"
	"go-my-life/internal/infrastructure"
	service "go-my-life/internal/service/money"
	"log"
	"mime/multipart"
	"strconv"
	"strings"
)

// MoneyMount 收入支出记账的web层
func MoneyMount() *fiber.App {
	app := fiber.New()

	app.Post("/transaction/list", func(c *fiber.Ctx) error {
		param := &moneydb.TransactionQueryParam{}
		// 解析JSON RequestBody
		err := c.BodyParser(param)
		if err != nil {
			return err
		}
		param.UserIds = append(param.UserIds, c.Locals("userId").(int64))
		// 获取结果
		res, total, err := service.QueryTransactions(*param)
		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"code":    "500",
				"error":   err.Error(),
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"data":    res,
			"total":   total,
		})
	})

	app.Post("/transaction/analysis", func(c *fiber.Ctx) error {
		param := &service.TransactionsAnalysisParam{}

		// 解析JSON RequestBody
		err := c.BodyParser(param)
		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
				"code":    "500",
			})
		}
		param.UserIds = append(param.UserIds, c.Locals("userId").(int64))
		results, err := service.TransactionsAnalysis(param)
		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
				"code":    "500",
			})
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"data":    results,
		})

	})

	app.Delete("/transaction", func(ctx *fiber.Ctx) error {
		id := ctx.QueryInt("id")
		if id == 0 {
			return ctx.JSON(&fiber.Map{
				"success": false,
			})
		}
		err := service.DeleteTransaction(int64(id))
		if err != nil {
			return ctx.JSON(&fiber.Map{
				"success": false,
				"code":    "500",
				"error":   err,
			})
		}
		return ctx.JSON(&fiber.Map{
			"success": true,
		})
	})

	// transaction put
	app.Put("/transaction", func(ctx *fiber.Ctx) error {
		transaction := &money.Transaction{}
		err := ctx.BodyParser(transaction)
		if err != nil {
			log.Println(err)
			return ctx.JSON(&fiber.Map{
				"success": false,
				"error":   "param parse failed!",
			})
		}
		err = service.PutTransaction(ctx.Locals("userId").(int64), transaction)
		success := true
		if err != nil {
			success = false
		}
		return ctx.JSON(&fiber.Map{
			"success": success,
			"error":   err,
		})
	})

	// batch put
	app.Put("/transaction/batch", func(ctx *fiber.Ctx) error {
		var transactions []*moneydb.Transaction
		err := ctx.BodyParser(&transactions)
		if err != nil {
			log.Println(err)
			return ctx.JSON(&fiber.Map{
				"success": false,
				"error":   "param parse failed!",
			})
		}

		// 获取用户id
		userId := ctx.Locals("userId").(int64)

		num, err := service.BatchPutTransaction(userId, transactions)
		if err != nil {
			return ctx.JSON(&fiber.Map{
				"success": false,
				"error":   err,
				"data":    num,
			})
		}
		return ctx.JSON(&fiber.Map{
			"success": true,
			"data":    num,
		})
	})

	app.Get("/transactionCategory", func(ctx *fiber.Ctx) error {
		res, err := service.AllCategory()
		if err != nil {
			return ctx.JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
				"code":    "500",
			})
		}
		return ctx.JSON(&fiber.Map{
			"success": true,
			"data":    res,
		})
	})

	app.Put("/transactionCategory", func(ctx *fiber.Ctx) error {
		param := &moneydb.TransactionCategory{}
		err := ctx.BodyParser(param)
		if err != nil {
			return ctx.JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
				"code":    "500",
			})
		}
		err = service.PutTransactionCategory(param)
		if err != nil {
			return ctx.JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
				"code":    "500",
			})
		}
		return ctx.JSON(&fiber.Map{
			"success": true,
		})
	})

	app.Get("/transactionAccount", func(ctx *fiber.Ctx) error {
		res, err := service.AllAccounts()
		if err != nil {
			return ctx.JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
				"code":    "500",
			})
		}
		return ctx.JSON(&fiber.Map{
			"success": true,
			"data":    res,
		})
	})

	app.Put("/transactionAccount", func(ctx *fiber.Ctx) error {
		param := &moneydb.TransactionAccount{}
		err := ctx.BodyParser(param)
		if err != nil {
			return err
		}
		res := service.PutTransactionAccount(param)
		return ctx.JSON(&fiber.Map{
			"success": true,
			"data":    res,
		})
	})

	// 解析Excel
	app.Post("/transaction/excel", func(ctx *fiber.Ctx) error {

		file, err := ctx.FormFile("file")
		if err != nil {
			print(err)
			return err
		}
		//log.Println(file)
		filename := file.Filename
		if !strings.HasSuffix(filename, ".csv") {
			return ctx.JSON(&fiber.Map{
				"success":  false,
				"errorMsg": "错误的文件格式",
			})
		}

		csvFile, err := file.Open()
		if err != nil {
			print(err)
			return err
		}

		// 延迟关闭文件
		defer func(csvFile multipart.File) {
			err := csvFile.Close()
			if err != nil {
				log.Print(err)
			}
		}(csvFile)

		transactions, err := service.ProcessTransactionExcel(ctx.Locals("userId").(int64), csvFile)

		return ctx.JSON(&fiber.Map{
			"success": true,
			"data":    transactions,
		})
	})

	app.Get("/transaction/category_analysis", func(ctx *fiber.Ctx) error {
		desc := ctx.Query("desc")

		res := service.AnalysisCategory(desc)
		return ctx.JSON(&fiber.Map{
			"success": true,
			"data":    res,
		})
	})

	app.Get("/transaction/category_immigrate", func(ctx *fiber.Ctx) error {
		param := moneydb.TransactionQueryParam{
			PageSize:  100,
			PageIndex: 1,
			StartTime: "2020-01-01",
			EndTime:   "2025-01-01",
		}

		transactions, _, _ := service.QueryTransactions(param)
		for transactions != nil && len(transactions) > 0 {
			for _, transaction := range transactions {
				doc := moneydb.TransactionIndex{
					Id:          transaction.Id,
					Description: transaction.Description,
					Type:        transaction.Type,
					Category:    transaction.Category,
				}
				infrastructure.EsClient.Index(moneydb.EsIndex).Id(strconv.FormatInt(transaction.Id, 10)).Request(doc).Do(context.Background())
			}
			param.PageIndex = param.PageIndex + 1
			transactions, _, _ = service.QueryTransactions(param)
		}
		return ctx.JSON(&fiber.Map{
			"success": true,
		})
	})
	return app
}
