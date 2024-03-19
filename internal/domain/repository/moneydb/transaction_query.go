package moneydb

import (
	"go-my-life/internal/infrastructure"
	"log"
	"time"
)

type TransactionQueryParam struct {
	UserIds    []int64 `json:"userIds"`
	Type       int     `json:"type"`
	StartTime  string  `json:"startTime"`
	EndTime    string  `json:"endTime"`
	PageSize   int     `json:"pageSize"`
	PageIndex  int     `json:"pageIndex"`
	AmountDesc bool    `json:"amountDesc"`
}

// QueryTransaction 复杂查询
func QueryTransaction(param TransactionQueryParam) ([]Transaction, int64) {
	session := infrastructure.Mysql.Where("is_deleted = 0")
	if param.Type != 0 {
		session = session.And("type = ?", param.Type)
	}
	if len(param.UserIds) > 0 {
		session = session.In("user_id", param.UserIds)
	}
	// 格式是 2023-09-12
	if param.StartTime != "" {
		transactionStartTime, _ := time.Parse("2006-01-02", param.StartTime)
		session = session.And("time >= ?", transactionStartTime)
	}
	if param.EndTime != "" {
		transactionEndTime, _ := time.Parse("2006-01-02", param.EndTime)
		transactionEndTime = transactionEndTime.Add(24 * time.Hour)
		// 记得加一天
		session = session.And("time < ?", transactionEndTime)
	}

	// 保存查询条件
	conditions := session.Conds()

	// 执行完会重置查询条件
	total, _ := session.Count(new(Transaction))

	session.And(conditions)
	if param.PageSize != 0 && param.PageIndex != 0 {
		session = session.Limit(param.PageSize, (param.PageIndex-1)*param.PageSize)
	}

	// 排序
	if param.AmountDesc {
		session.Desc("amount")
	} else {
		session.Desc("time")
	}

	var moneys []Transaction
	err := session.Find(&moneys)
	if err != nil {
		log.Println(err)
		return moneys, total
	}
	return moneys, total
}
