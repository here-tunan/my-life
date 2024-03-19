package userdb

import (
	"errors"
	"go-my-life/internal/infrastructure"
	"go-my-life/pkg/model"
)

type User struct {
	// Id
	Id int64 `json:"id"`
	// 账户
	Account string `json:"account"`
	// 密码
	Password string `json:"password"`
	// 名称
	Name string `json:"name"`
	// 描述
	Desc string `json:"desc"`
	// 头像
	Avatar string `json:"avatar"`
	// 额外信息
	Extra string `json:"extra"`
	// 是否删除
	IsDeleted bool `json:"isDeleted"`
	// 系统创建时间
	GmtCreate model.LocalTime `json:"gmtCreate" xorm:"updated"`
	// 系统更新时间
	GmtModified model.LocalTime `json:"gmtModified" xorm:"updated"`
}

func (user *User) GetUser() error {
	session := infrastructure.Mysql.Where("is_deleted = ?", 0)

	if user.Id != 0 {
		session.And("id = ?", user.Id)
	}

	if user.Account != "" {
		session.And("account = ?", user.Account)
	}

	if user.Password != "" {
		session.And("password = ?", user.Password)
	}

	session.Limit(1)

	has, err := session.Get(user)

	if !has {
		return errors.New("no this user")
	}
	return err
}

func (user *User) PutUser() error {
	if user.Id == 0 {
		// 新增
		_, err := infrastructure.Mysql.InsertOne(user)
		if err != nil {
			return err
		}
	} else {
		// 更新（默认 只更新non-empty 和 non-zero的字段）
		_, err := infrastructure.Mysql.ID(user.Id).Omit("GmtCreate").Update(user)
		if err != nil {
			return err
		}
	}
	return nil
}
