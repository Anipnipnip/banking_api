package service

import (
	"banking-api/internal/model"
	"banking-api/internal/repository"
)

func CreateUser(user model.User) model.User {

	user = repository.CreateUser(user)

	return user
}
