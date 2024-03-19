package service

import (
	"go-my-life/internal/domain/entity/user"
	"go-my-life/internal/domain/repository/userdb"
)

func GetUser(user *userdb.User) (*userdb.User, error) {
	err := user.GetUser()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func PutUser(user *user.User) error {
	userDb := &userdb.User{
		Id:       user.Id,
		Account:  user.Account,
		Password: user.Password,
		Name:     user.Name,
		Desc:     user.Desc,
		Avatar:   user.Avatar,
		Extra:    user.Extra,
	}

	err := userDb.PutUser()
	user.Id = userDb.Id
	if err != nil {
		return err
	}
	return nil
}
