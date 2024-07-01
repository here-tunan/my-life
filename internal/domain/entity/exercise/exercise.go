package exercise

import (
	"go-my-life/internal/domain/repository/exercisedb"
	"go-my-life/pkg/model"
)

type Log struct {
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
}

func (exercise *Log) Create() error {
	record := exercisedb.ExerciseLog{
		Date:        exercise.Date,
		Description: exercise.Description,
		UserId:      exercise.UserId,
		Tag:         exercise.Tag,
		Duration:    exercise.Duration,
	}
	return record.Insert()
}
func (exercise *Log) Update() error {
	record := exercisedb.ExerciseLog{
		Id:          exercise.Id,
		Duration:    exercise.Duration,
		Description: exercise.Description,
		UserId:      exercise.UserId,
		Tag:         exercise.Tag,
		Date:        exercise.Date,
	}
	return record.Update()
}

func (exercise *Log) Delete() error {
	record := exercisedb.ExerciseLog{
		Id: exercise.Id,
	}
	return record.Delete()
}
