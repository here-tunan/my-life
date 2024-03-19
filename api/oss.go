package api

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Policy struct {
	Expiration string `json:"expiration"`
	Conditions []any  `json:"conditions"`
}

func OssMount() *fiber.App {
	app := fiber.New()

	app.Get("", func(ctx *fiber.Ctx) error {
		now := time.Now().Unix()
		expiredEnd := now + 3600
		var tokenExpire = time.Unix(expiredEnd, 0).UTC().Format("2006-01-2T15:04:05Z")

		// https://help.aliyun.com/zh/oss/developer-reference/postobject#section-d5z-1ww-wdb 注意格式，
		// 要不然生成的policy在传给阿里云的时候很容易出现 Invalid Json 的报错
		policy := Policy{
			Expiration: tokenExpire,
			Conditions: []any{
				map[string]string{"bucket": "your-bucket"},
				[]any{"content-length-range", 0, 1048576000},
				[]string{"eq", "$success_action_status", "200"},
				//[]string{"starts-with", "$key", "user/eric/"},
				//[]any{"in", "$content-type", []string{"image/jpg", "image/png"}},
				//[]any{"not-in", "$cache-control", []string{"no-cache"}},
			},
		}

		policyBytes, _ := json.Marshal(policy)
		policyText := base64.StdEncoding.EncodeToString(policyBytes)

		accessKey := "your key" // 你的AccessKey（secret）
		key := []byte(accessKey)
		// 设置key（）
		hmacBytes := hmac.New(sha1.New, key)
		// 设置hmac内容
		hmacBytes.Write([]byte(policyText))
		signature := base64.StdEncoding.EncodeToString(hmacBytes.Sum(nil))
		fmt.Println("Signature:", signature)

		oss := fiber.Map{
			"accessId":  "your accessId",
			"host":      "your host",
			"policy":    policyText,
			"signature": signature,
			"expire":    expiredEnd,
			"dir":       "your-dirs/",
		}

		return ctx.JSON(&fiber.Map{
			"success": true,
			"data":    oss,
		})
	})

	return app
}
