package familydb

import (
	"errors"
	"go-my-life/internal/infrastructure"
	"go-my-life/pkg/model"
	"go-my-life/pkg/myerror"
)

// Family 对应数据库实体
type Family struct {
	// Id
	Id int64 `json:"id"`
	// 名称
	Name string `json:"amount"`
	// 描述
	Desc string `json:"desc"`
	// 头像
	Avatar string `json:"avatar"`
	// 是否删除
	IsDeleted bool `json:"isDeleted"`
	// 系统创建时间
	GmtCreate model.LocalTime `json:"gmtCreate" xorm:"updated"`
	// 系统更新时间
	GmtModified model.LocalTime `json:"gmtModified" xorm:"updated"`
}

func (family *Family) Insert() error {
	if family.Name == "" {
		return errors.New("family name can't be blank")
	}
	// insert family
	_, err := infrastructure.Mysql.InsertOne(family)

	return err
}

func (family *Family) Update() error {
	if family.Name == "" {
		return errors.New("family name can't be blank")
	}
	if family.Id == 0 {
		return errors.New("family id can't be null")
	}
	// update
	_, err := infrastructure.Mysql.ID(family.Id).Omit("GmtCreated").Update(family)
	return err
}

func (family *Family) Get() error {
	session := infrastructure.Mysql.Where("is_deleted = 0")

	if family.Id != 0 {
		session.ID(family.Id)
	}

	session.Limit(1)
	has, err := session.Get(family)
	if !has {
		err = myerror.NewError(myerror.NoFamily)
	}
	if err != nil {
		return err
	}
	return nil
}
