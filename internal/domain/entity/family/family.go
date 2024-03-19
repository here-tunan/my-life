package family

import (
	"errors"
	"go-my-life/internal/domain/repository/familydb"
)

// Family 领域实体 聚合根（family 信息，家庭成员……）
type Family struct {
	// Id
	Id int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 描述
	Desc string `json:"desc"`
	// 头像
	Avatar string `json:"avatar"`
	// 家庭成员
	Members []*Member `json:"members"`
}

func (family *Family) GetFamily() error {
	record := familydb.Family{
		Id: family.Id,
	}
	err := record.Get()
	if err != nil {
		return err
	}
	family.Id = record.Id
	family.Name = record.Name
	family.Desc = record.Desc
	family.Avatar = record.Avatar

	members, err := familydb.QueryAllMember(family.Id)
	if err != nil {
		return err
	}
	if len(members) != 0 {
		for _, member := range members {
			m := &Member{
				Id:     member.Id,
				UserId: member.UserId,
			}
			err = m.GetByUserId()
			if err != nil {
				return err
			}
			family.Members = append(family.Members, m)
		}
	}

	return nil
}

// CreateFamily 创建一个家庭
func (family *Family) CreateFamily() error {
	if len(family.Members) == 0 {
		return errors.New("no first member to init a family")
	}

	record := familydb.Family{
		Name:   family.Name,
		Desc:   family.Desc,
		Avatar: family.Avatar,
	}
	err := record.Insert()
	if err != nil {
		return err
	}

	// addMember
	member := family.Members[0]
	member.FamilyId = record.Id
	err = member.CreateMember()
	if err != nil {
		return err
	}
	family.Id = record.Id
	return err
}

func (family *Family) UpdateFamily() error {
	record := familydb.Family{
		Id:     family.Id,
		Name:   family.Name,
		Desc:   family.Desc,
		Avatar: family.Avatar,
	}
	err := record.Update()
	family.Id = record.Id
	return err
}
