package service

import (
	"go-my-life/internal/domain/entity/money"
	"go-my-life/internal/domain/repository/moneydb"
	timeutil "go-my-life/pkg/utils"
)

type TransactionQueryResponse struct {
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
	Time string `json:"time"`
}

func DeleteTransaction(id int64) error {
	transaction := money.Transaction{
		Id: id,
	}
	return transaction.Delete()
}

func PutTransaction(userId int64, transaction *money.Transaction) error {
	transaction.UserId = userId
	if transaction.Id == 0 {
		return transaction.Create()
	} else {
		return transaction.Update()
	}
}

// BatchPutTransaction 批量操作内含事务管理
func BatchPutTransaction(userId int64, transactions []*moneydb.Transaction) (int, error) {
	for _, transaction := range transactions {
		transaction.UserId = userId
	}
	err := moneydb.BatchPutTransaction(transactions)
	if err != nil {
		return 0, err
	}
	return len(transactions), nil
}

func QueryTransactions(param moneydb.TransactionQueryParam) ([]TransactionQueryResponse, int64, error) {
	transactionList, total := moneydb.QueryTransaction(param)
	if len(transactionList) == 0 {
		return nil, 0, nil
	}
	var resList []TransactionQueryResponse
	for _, transaction := range transactionList {
		var res TransactionQueryResponse
		res.Id = transaction.Id
		res.Type = transaction.Type
		res.UserId = transaction.UserId
		res.Account = transaction.Account
		res.Category = transaction.Category
		res.Amount = transaction.Amount
		res.Description = transaction.Description
		res.Time = timeutil.LocalTimeFormat(transaction.Time)

		resList = append(resList, res)
	}

	return resList, total, nil
}
