package familydb

import (
	"errors"
	"go-my-life/internal/infrastructure"
	"go-my-life/pkg/model"
	"go-my-life/pkg/myerror"
)

type FamilyMember struct {
	// Id
	Id int64 `json:"id"`
	// family id
	FamilyId int64 `json:"familyId"`
	// user id
	UserId int64 `json:"userId"`
	// 是否是创建者
	IsCreator bool `json:"isCreator"`
	// 是否删除
	IsDeleted bool `json:"isDeleted"`
	// 系统创建时间
	GmtCreate model.LocalTime `json:"gmtCreate" xorm:"updated"`
	// 系统更新时间
	GmtModified model.LocalTime `json:"gmtModified" xorm:"updated"`
}

// Insert 新增一个家庭成员
func (member *FamilyMember) Insert() error {
	if member.UserId == 0 {
		return errors.New("userId can not be null")
	}
	if member.FamilyId == 0 {
		return errors.New("familyId can not be null")
	}

	_, err := infrastructure.Mysql.Insert(member)
	return err
}

func (member *FamilyMember) Get() error {
	has, err := infrastructure.Mysql.Where("is_deleted = 0").Get(member)
	if !has {
		err = myerror.NewError(myerror.NotFamilyMember)
	}
	if err != nil {
		return err
	}
	return nil
}

func QueryAllMember(familyId int64) ([]FamilyMember, error) {
	var res []FamilyMember
	err := infrastructure.Mysql.Where("is_deleted = 0").And("family_id = ?", familyId).Find(&res)
	return res, err
}
