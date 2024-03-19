package money

import (
	"go-my-life/internal/domain/repository/moneydb"
	"go-my-life/pkg/model"
)

type Transaction struct {
	// Id
	Id int64 `json:"id"`
	// 金额
	Amount float64 `json:"amount"`
	// 描述
	Description string `json:"description"`
	// 使用人
	UserId int64 `json:"userId"`
	// 金额类型（收入/支出）
	Type int `json:"type"`
	// 用途分类
	Category int `json:"category"`
	// 账户
	Account int `json:"account"`
	// 操作时间
	Time model.LocalTime `json:"time"`
}

func (transaction *Transaction) Create() error {
	record := moneydb.Transaction{
		Amount:      transaction.Amount,
		Description: transaction.Description,
		UserId:      transaction.UserId,
		Type:        transaction.Type,
		Category:    transaction.Category,
		Account:     transaction.Account,
		Time:        transaction.Time,
	}
	return record.Insert()
}

func (transaction *Transaction) Update() error {
	record := moneydb.Transaction{
		Id:          transaction.Id,
		Amount:      transaction.Amount,
		Description: transaction.Description,
		UserId:      transaction.UserId,
		Type:        transaction.Type,
		Category:    transaction.Category,
		Account:     transaction.Account,
		Time:        transaction.Time,
	}
	return record.Update()
}

func (transaction *Transaction) Delete() error {
	record := moneydb.Transaction{
		Id: transaction.Id,
	}
	return record.Delete()
}
