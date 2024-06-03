package impl

import (
	"github.com/hrishikeshgoyal/quest/go-microservice-fk/entity"
	"github.com/hrishikeshgoyal/quest/go-microservice-fk/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepo,
	}
}

func (u UserService) GetUser(id string) *entity.User {
	return u.userRepository.GetUser(id)
}

func (u UserService) GetUserByName(name string) *entity.User {
	return u.userRepository.GetUserByName(name)
}

func (u UserService) CreateUser(user *entity.User) *entity.User {
	return u.userRepository.CreateUser(user)
}

func (u UserService) GetAllUsers() []entity.User {
	return u.userRepository.GetAll()
}
