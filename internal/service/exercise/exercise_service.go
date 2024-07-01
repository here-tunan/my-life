package exercise

import (
	"go-my-life/internal/domain/entity/exercise"
	"go-my-life/internal/domain/repository/exercisedb"
)

func PutExercise(userId int64, exercise *exercise.Log) error {
	exercise.UserId = userId
	if exercise.Id == 0 {
		return exercise.Create()
	} else {
		return exercise.Update()
	}
}

func ListExerciseLogs(userId int64, param exercisedb.ExerciseQuery) []exercisedb.ExerciseLog {
	param.UserId = userId
	return exercisedb.Query(param)
}
