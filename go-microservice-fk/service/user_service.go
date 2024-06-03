package service

import (
	"github.com/hrishikeshgoyal/quest/go-microservice-fk/entity"
)

type UserService interface {
	GetUser(id string) *entity.User

	GetUserByName(name string) *entity.User
	CreateUser(user *entity.User) *entity.User
	GetAllUsers() []entity.User
}
