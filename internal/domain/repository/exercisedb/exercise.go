package exercisedb

import (
	"errors"
	"go-my-life/internal/infrastructure"
	"go-my-life/pkg/model"
	"log"
	"time"
)

type ExerciseLog struct {
	// Id
	Id int64 `json:"id"`
	// 用户ID
	UserId int64 `json:"userId"`
	// 运动类型
	Tag string `json:"tag"`
	// 描述
	Description string `json:"description"`
	// 运动时长
	Duration int `json:"duration"`
	// 运动日期
	Date model.LocalDate `json:"date"`
	// 是否删除
	IsDeleted bool `json:"isDeleted"`
	// 系统创建时间
	GmtCreate model.LocalTime `json:"gmtCreate" xorm:"updated"`
	// 系统更新时间
	GmtModified model.LocalTime `json:"gmtModified" xorm:"updated"`
}

type ExerciseQuery struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	UserId    int64  `json:"userId"`
}

func (record *ExerciseLog) Insert() error {
	// 插入mysql
	_, err := infrastructure.Mysql.InsertOne(record)
	if err != nil {
		return err
	}

	return err
}

func (record *ExerciseLog) Update() error {
	_, err := infrastructure.Mysql.ID(record.Id).Omit("GmtCreate").Update(record)
	if err != nil {
		return err
	}
	return nil
}

func (record *ExerciseLog) Delete() error {
	if record.Id == 0 {
		return errors.New("exercise id can't be null")
	}
	_, err := infrastructure.Mysql.Table(&ExerciseLog{}).ID(record.Id).Update(map[string]interface{}{"is_deleted": true})
	if err != nil {
		return err
	}
	return err
}

func Query(param ExerciseQuery) []ExerciseLog {
	session := infrastructure.Mysql.Where("is_deleted = 0")
	if param.StartDate != "" {
		startDate, _ := time.Parse("2006-01-02", param.StartDate)
		session = session.And("date >= ?", startDate)
	}
	if param.EndDate != "" {
		EndDate, _ := time.Parse("2006-01-02", param.EndDate)
		EndDate = EndDate.Add(24 * time.Hour)
		// 记得加一天
		session = session.And("date < ?", EndDate)
	}

	var logs []ExerciseLog
	err := session.Find(&logs)
	if err != nil {
		log.Println(err)
		return logs
	}
	return logs
}
