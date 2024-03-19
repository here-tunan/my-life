package moneydb

import (
	"go-my-life/internal/infrastructure"
	"time"
)

type TransactionAccount struct {
	// id
	Id int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 描述
	Desc string `json:"desc"`
	// 是否删除
	IsDeleted bool `json:"isDeleted"`
	// 系统创建时间
	GmtCreate time.Time `json:"gmtCreate" xorm:"updated"`
	// 系统更新时间
	GmtModified time.Time `json:"gmtModified" xorm:"updated"`
}

func AllAccounts() ([]TransactionAccount, error) {
	session := infrastructure.Mysql.Where("is_deleted = 0")
	var accounts []TransactionAccount
	err := session.Find(&accounts)
	if err != nil {
		return nil, err
	}
	return accounts, err
}

func (account *TransactionAccount) Insert() error {
	_, err := infrastructure.Mysql.InsertOne(account)
	return err
}

func (account *TransactionAccount) Update() error {
	_, err := infrastructure.Mysql.ID(account.Id).Omit("GmtCreate").Update(account)
	return err
}
