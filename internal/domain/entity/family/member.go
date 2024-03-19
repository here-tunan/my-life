package family

import (
	"go-my-life/internal/domain/repository/familydb"
	"go-my-life/internal/domain/repository/userdb"
	"go-my-life/pkg/myerror"
)

// Member 家庭成员实体
type Member struct {
	Id int64 `json:"id"`
	// family id
	FamilyId int64 `json:"familyId"`
	// user id
	UserId int64 `json:"userId"`
	// 是否是创建者
	IsCreator bool `json:"isCreator"`
	// 名称
	Name string `json:"name"`
	// 描述
	Desc string `json:"desc"`
	// 头像
	Avatar string `json:"avatar"`
}

func (member *Member) GetByUserId() error {
	if member.UserId == 0 {
		return myerror.NewError(myerror.WrongParam)
	}
	// 查询user信息
	user := &userdb.User{
		Id: member.UserId,
	}
	err := user.GetUser()
	if err != nil {
		return err
	}
	memberRecord := &familydb.FamilyMember{
		UserId: member.UserId,
	}
	err = memberRecord.Get()
	if err != nil {
		return err
	}

	member.Id = memberRecord.Id
	member.IsCreator = memberRecord.IsCreator
	member.FamilyId = memberRecord.FamilyId
	member.Name = user.Name
	member.Desc = user.Desc
	member.Avatar = user.Avatar

	return nil
}

// CreateMember 增加家庭成员
func (member *Member) CreateMember() error {
	if member.UserId == 0 {
		return myerror.NewError(myerror.WrongParam)
	}

	// 查询user信息
	user := &userdb.User{
		Id: member.UserId,
	}
	err := user.GetUser()
	if err != nil {
		return err
	}
	record := &familydb.FamilyMember{
		FamilyId:  member.FamilyId,
		UserId:    member.UserId,
		IsCreator: member.IsCreator,
	}

	// 新增成员
	err = record.Insert()
	if err != nil {
		return err
	}
	member.Id = record.Id
	member.Avatar = user.Avatar
	member.Desc = user.Desc
	member.Name = user.Name
	return nil
}
