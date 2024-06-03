package repository

//go:generate mockgen -source=user_repository.go -destination=mocks/user_repository.go

import "github.com/hrishikeshgoyal/quest/go-microservice-fk/entity"

type UserRepository interface {
	GetUser(id string) *entity.User
	GetUserByName(name string) *entity.User
	CreateUser(user *entity.User) *entity.User
	GetAll() []entity.User
}
