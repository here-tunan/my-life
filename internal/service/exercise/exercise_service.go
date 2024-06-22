package exercise

import (
	"go-my-life/internal/domain/entity/exercise"
)

func PutExercise(userId int64, exercise *exercise.Log) error {
	exercise.UserId = userId
	if exercise.Id == 0 {
		return exercise.Create()
	} else {
		return exercise.Update()
	}
}
