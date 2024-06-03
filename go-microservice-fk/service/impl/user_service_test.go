package impl

import (
	"github.com/golang/mock/gomock"
	"github.com/hrishikeshgoyal/quest/go-microservice-fk/entity"
	mock_repository "github.com/hrishikeshgoyal/quest/go-microservice-fk/repository/mocks"
	"testing"
)

func TestUserService_GetUser(t *testing.T) {
	cases := []struct {
		name string
		id   string
		user *entity.User
	}{
		{
			name: "happyCase",
			id:   "123",
			user: &entity.User{
				Name:  "Hrishi",
				Email: "hrishi@gmail.com",
			},
		},
	}

	for _, c := range cases {
		ctrl := gomock.NewController(t)

		// 👇 create new user repo
		mockUserRepo := mock_repository.NewMockUserRepository(ctrl)

		us := NewUserService(mockUserRepo)

		mockUserRepo.EXPECT().GetUser(gomock.Eq("123")).Return(c.user)
		returnedUser := us.GetUser(c.id)

		if returnedUser != c.user {
			t.Fatalf("expected %v, got: %v", c.user, returnedUser)
		}
	}

}
