package service

import (
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"
	port "rpolnx.com.br/crud-sql/internal/crud-sql/domain/ports"
)

type userService struct {
	userPort port.UserPort
}

func (svc *userService) GetAllUsers() ([]*model.User, error) {
	return svc.userPort.GetAllUsersOut()
}

func (svc *userService) GetOneUser(id *int64) (*model.User, error) {
	return svc.userPort.GetOneUserOut(id)
}

func (svc *userService) CreateUser(user *model.User) (*int64, error) {
	return svc.userPort.CreateUserOut(user)
}

func (svc *userService) UpdateUser(id *int64, user *model.User) (*int64, error) {
	return svc.userPort.UpdateUserOut(id, user)
}

func (svc *userService) DeleteUser(id *int64) error {
	return svc.userPort.DeleteUserOut(id)
}

func NewUserService(userPort port.UserPort) port.UserUseCase {
	return &userService{
		userPort: userPort,
	}
}
