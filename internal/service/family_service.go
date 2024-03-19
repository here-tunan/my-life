package service

import (
	"go-my-life/internal/domain/entity/family"
)

func GetFamily(userId int64) (*family.Family, error) {
	// 寻找家庭Id
	memberEntity := &family.Member{
		UserId: userId,
	}
	err := memberEntity.GetByUserId()
	if err != nil {
		return nil, err
	}
	familyEntity := &family.Family{
		Id: memberEntity.FamilyId,
	}
	err = familyEntity.GetFamily()
	return familyEntity, err
}

// CreateFamily 创建一个新的家庭
func CreateFamily(userId int64, familyEntity *family.Family) error {
	member := &family.Member{
		UserId:    userId,
		IsCreator: true,
	}
	// 家庭成员
	familyEntity.Members = append(familyEntity.Members, member)

	err := familyEntity.CreateFamily()
	return err
}
