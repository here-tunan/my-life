package service

import (
	"encoding/csv"
	"fmt"
	"go-my-life/internal/domain/repository/moneydb"
	"go-my-life/pkg/model"
	"log"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

func ProcessTransactionExcel(userId int64, csvFile multipart.File) ([]moneydb.Transaction, error) {

	// 创建 CSV Reader
	reader := csv.NewReader(csvFile)
	// 这些csv文件都不规范，行与行之间的fields num不一样，会导致错误，加上这个就好了
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true // 忽略未被双引号括起来的字段

	// 读取 CSV 数据
	records, err := reader.ReadAll()
	if err != nil {
		log.Print("Failed to read transaction CSV record:", err)
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}

	// ZWNBSP 微信的第一个单元格里面有这么个玩意，没法进行匹配, 用第二行进行匹配
	if strings.HasPrefix(records[1][0], "微信昵称") {
		fmt.Print("处理微信账单!")
		fmt.Println(records[1][0])
		return processWeChatBill(userId, records), nil
	}

	if strings.Contains(records[3][0], "支付宝账户") {
		fmt.Print("处理支付宝账单!")
		after, _ := strings.CutPrefix(records[2][0], "姓名：")
		fmt.Printf("支付宝用户：%s\n", after)
		return processAlipayBill(userId, records), nil
	}

	return processMyExcel(userId, records), nil
}

func processWeChatBill(userId int64, records [][]string) []moneydb.Transaction {
	var transactions []moneydb.Transaction
	isBillStart := false
	for _, record := range records {
		if isBillStart {
			// 开始获取真正的账单行数据
			// 第一列：交易时间
			timeStr := record[0]
			localTime, _ := time.Parse("2006-01-02 15:04:05", timeStr)
			// 第三列：交易对方; 第四列：商品
			description := record[2] + "_" + record[3]
			// 第五列：收入/支出
			typeId := 1
			if record[4] == "收入" {
				typeId = 1
			} else {
				typeId = 2
			}

			// 先暂且不管
			category := AnalysisCategory(description)

			// 第六列
			amount, _ := strconv.ParseFloat(strings.Trim(record[5], "¥ "), 64)
			transaction := moneydb.Transaction{
				Amount:      amount,
				Description: description,
				UserId:      userId,
				Type:        typeId,
				Category:    category,
				Account:     1,
				Time:        model.LocalTime(localTime),
			}
			transactions = append(transactions, transaction)
			continue
		}
		if record[0] == "交易时间" {
			isBillStart = true
		}
	}
	return transactions
}

func processAlipayBill(userId int64, records [][]string) []moneydb.Transaction {
	var transactions []moneydb.Transaction
	isBillStart := false
	for _, record := range records {
		if isBillStart {
			// 开始获取真正的账单行数据
			// 第二列 交易分类（筛选掉投资理财的）
			if record[1] == "投资理财" {
				continue
			}

			// 第一列：交易时间
			timeStr := record[0]
			localTime, _ := time.Parse("2006/1/2 15:04", timeStr)

			// 第三列：交易对方; 第五列：商品
			description := record[2] + "_" + record[4]

			// 先暂且不管
			category := AnalysisCategory(description)

			// 第六列：收入/支出
			typeId := 1
			if record[5] == "收入" {
				typeId = 1
			} else {
				typeId = 2
			}
			// 第七列 金额
			amount, _ := strconv.ParseFloat(record[6], 64)

			// 第八列收付款方式 有一个要特殊处理一下
			account := 2
			if strings.Contains(record[7], "关爱通") {
				account = 12 // 关爱通
			}
			transaction := moneydb.Transaction{
				Amount:      amount,
				Description: description,
				UserId:      userId,
				Type:        typeId,
				Category:    category,
				Account:     account,
				Time:        model.LocalTime(localTime),
			}
			transactions = append(transactions, transaction)
			continue
		}
		if record[0] == "交易时间" {
			isBillStart = true
		}
	}
	return transactions
}

func processMyExcel(userId int64, records [][]string) []moneydb.Transaction {
	var transactions []moneydb.Transaction
	for _, record := range records {
		if record[6] == "是" {
			continue
		}

		// 第1列 金额钱
		amount, _ := strconv.ParseFloat(record[0], 64)
		// 第2列 描述
		description := record[1]
		// 第3列 人
		//var userId int64 = 1
		//if record[2] == "Z" {
		//	userId = 2
		//}

		accountMap := map[string]int{
			"微信":   1,
			"支付宝":  2,
			"建设银行": 8,
			"中国银行": 6,
			"工商银行": 7,
			"农业银行": 9,
			"招商银行": 10,
			"医保":   5,
			"美团":   4,
			"给到":   12,
		}

		// 第四列 支出账户
		account := accountMap[record[3]]

		// 第5列 时间
		localTime, _ := time.Parse("2006-01-02", record[4])

		// 第6列 收支
		typeId := 1
		if record[5] == "收入" {
			typeId = 1
		} else {
			typeId = 2
		}

		// 第8列 类型
		categoryMap := map[string]int{
			"饮食":        1,
			"医疗":        3,
			"聚会":        5,
			"购物（含生活用品）": 2,
			"电子产品":      9,
			"房租":        10,
			"其他":        101,
			"交通":        11,
			"通讯":        8,
			"运动健身":      4,
			"教育":        6,
			"娱乐":        5,
			"水电煤气费":     7,
			"收入":        14,
		}
		category := categoryMap[record[7]]

		transaction := moneydb.Transaction{
			Amount:      amount,
			Description: description,
			UserId:      userId,
			Type:        typeId,
			Category:    category,
			Account:     account,
			Time:        model.LocalTime(localTime),
		}
		transactions = append(transactions, transaction)
		continue
	}
	return transactions
}
