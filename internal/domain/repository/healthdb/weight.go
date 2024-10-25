package healthdb

import (
	"go-my-life/internal/infrastructure"
	"go-my-life/pkg/model"
	"log"
	"time"
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
	// 是否删除
	IsDeleted bool `json:"isDeleted"`
	// 系统创建时间
	GmtCreate model.LocalTime `json:"gmtCreate" xorm:"updated"`
	// 系统更新时间
	GmtModified model.LocalTime `json:"gmtModified" xorm:"updated"`
}

type WeightQueryParam struct {
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
	UserId    int64  `json:"userId"`
	Day       string `json:"day"`
	StartDay  string `json:"startDay"`
	EndDay    string `json:"endDay"`
}

func (weight *Weight) Insert() error {
	// 插入mysql
	_, err := infrastructure.Mysql.InsertOne(weight)
	if err != nil {
		return err
	}
	return nil
}

func (weight *Weight) Update() error {
	// 更新mysql
	_, err := infrastructure.Mysql.ID(weight.Id).Omit("GmtCreate").Update(weight)
	if err != nil {
		return err
	}

	return nil
}

func (weight *Weight) Delete() error {
	// 删除mysql
	_, err := infrastructure.Mysql.Table(&Weight{}).ID(weight.Id).Update(map[string]interface{}{"is_deleted": true})
	if err != nil {
		return err
	}
	return nil
}

func QueryWeights(param WeightQueryParam) ([]Weight, int64) {
	// 查询mysql
	session := infrastructure.Mysql.Where("is_deleted = 0")
	if param.UserId > 0 {
		session = session.And("user_id = ?", param.UserId)
	}
	// 格式是 2023-09-12
	if param.StartDay != "" {
		queryStartDay, _ := time.Parse("2006-01-02", param.StartDay)
		session = session.And("day >= ?", queryStartDay)
	}
	if param.EndDay != "" {
		queryEndDay, _ := time.Parse("2006-01-02", param.EndDay)
		queryEndDay = queryEndDay.Add(24 * time.Hour)
		// 记得加一天
		session = session.And("day < ?", queryEndDay)
	}
	if param.Day != "" {
		queryDay, _ := time.Parse("2006-01-02", param.Day)
		session = session.And("day = ?", queryDay)
	}

	// 保存查询条件
	conditions := session.Conds()

	// 执行完会重置查询条件
	total, _ := session.Count(new(Weight))

	session.And(conditions)
	if param.PageSize != 0 && param.PageIndex != 0 {
		session = session.Limit(param.PageSize, (param.PageIndex-1)*param.PageSize)
	}

	// 默认根据时间倒序排
	session.Desc("day")

	var records []Weight
	err := session.Find(&records)
	if err != nil {
		log.Println(err)
		return records, total
	}
	return records, total
}
