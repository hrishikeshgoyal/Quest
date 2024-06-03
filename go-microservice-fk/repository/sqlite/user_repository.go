package sqlite

import (
	"fmt"
	"github.com/hrishikeshgoyal/quest/go-microservice-fk/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB  *gorm.DB
	Url string
}

func NewUserRepository(url string) *UserRepository {
	DB, err := gorm.Open(sqlite.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&entity.User{})

	return &UserRepository{
		DB:  DB,
		Url: url,
	}
}

func (u UserRepository) GetUser(id string) *entity.User {
	user := entity.User{}
	u.DB.First(&user, id)
	return &user
}

func (u UserRepository) CreateUser(user *entity.User) *entity.User {
	u.DB.Create(user)
	return user
}

func (u UserRepository) GetAll() []entity.User {
	var users []entity.User
	u.DB.Find(&users)
	return users
}
