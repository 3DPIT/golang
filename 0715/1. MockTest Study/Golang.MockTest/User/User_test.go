package User_test

import (
	"testing"

	"fileMonitoring/Mirero.Golang.MockTest/User"
	"fileMonitoring/Mirero.Golang.MockTest/mocks"

	"github.com/golang/mock/gomock"
)

func TestUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUser := mocks.NewMockIUserInterface(mockCtrl)
	testUser := &User.User{IUser: mockUser}
	mockUser.EXPECT().AddUser(1, "sample test").Return(nil).Times(1)
	testUser.Use()

}
