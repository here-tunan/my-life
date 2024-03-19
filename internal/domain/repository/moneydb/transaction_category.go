package moneydb

import (
	"go-my-life/internal/infrastructure"
	"time"
)

type TransactionCategory struct {
	// id
	Id int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 收入/支出
	Type int `json:"type"`
	// 描述
	Desc string `json:"desc"`
	// 是否删除
	IsDeleted bool `json:"isDeleted"`
	// 系统创建时间
	GmtCreate time.Time `json:"gmtCreate" xorm:"updated"`
	// 系统更新时间
	GmtModified time.Time `json:"gmtModified" xorm:"updated"`
}

func AllCategory() ([]TransactionCategory, error) {
	session := infrastructure.Mysql.Where("is_deleted = 0")
	var categories []TransactionCategory
	err := session.Find(&categories)
	if err != nil {
		return nil, err
	}
	return categories, err
}

func (category *TransactionCategory) Insert() error {
	_, err := infrastructure.Mysql.InsertOne(category)
	return err
}

func (category *TransactionCategory) Update() error {
	_, err := infrastructure.Mysql.ID(category.Id).Omit("GmtCreate").Update(category)
	return err
}
