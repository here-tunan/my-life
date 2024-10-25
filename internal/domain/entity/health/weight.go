package health

import (
	"go-my-life/internal/domain/repository/healthdb"
	"go-my-life/pkg/model"
)

type Weight struct {
	// Id
	Id int64 `json:"id"`
	// 体重 KG
	Weight float64 `json:"weight"`
	// 用户id
	UserId int64 `json:"userId"`
	// 记录时间
	Day model.LocalDate `json:"day"`
}

func (weight *Weight) Put() error {
	// 先根据Day和UserId查询
	oldRecords, total := healthdb.QueryWeights(healthdb.WeightQueryParam{
		Day:    weight.Day.String(),
		UserId: weight.UserId,
	})

	if total > 0 {
		// 更新
		oldRecords[0].Weight = weight.Weight
		return oldRecords[0].Update()
	} else {
		// 新增
		newWeight := &healthdb.Weight{
			Weight: weight.Weight,
			UserId: weight.UserId,
			Day:    weight.Day,
		}
		return newWeight.Insert()
	}
}
