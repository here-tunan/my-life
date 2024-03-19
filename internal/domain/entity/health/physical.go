package health

import "go-my-life/pkg/model"

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
}
