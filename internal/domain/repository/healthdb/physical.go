package healthdb

import (
	"go-my-life/internal/infrastructure"
	"go-my-life/pkg/model"
	"log"
)

// Physical 身体状况
type Physical struct {
	// id
	Id int64 `json:"id"`
	// userId
	UserId int64 `json:"userId"`
	// 身高
	Height int `json:"height"`
	// 体重
	Weight int `json:"weight"`
	// 描述
	Desc string `json:"desc"`
	// 记录时间
	RecordTime model.LocalTime `json:"recordTime"`
	// 是否删除
	IsDeleted bool `json:"isDeleted"`
	// 系统创建时间
	GmtCreate model.LocalTime `json:"gmtCreate"`
	// 系统更新时间
	GmtModified model.LocalTime `json:"gmtModified"`
}

// QueryPhysicals 分页查询
func QueryPhysicals(userId int64, pageSize int, pageIndex int) ([]Physical, int64) {
	session := infrastructure.Mysql.Where("is_deleted = 0")
	if userId != 0 {
		session.And("user_id = ?", userId)
	}

	if pageSize != 0 && pageIndex != 0 {
		session = session.Limit(pageSize, (pageIndex-1)*pageSize)
	}

	var physicals []Physical
	err := session.Find(&physicals)
	if err != nil {
		log.Println(err)
		return physicals, 0
	}
	return physicals, 0
}
