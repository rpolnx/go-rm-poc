package port

import "rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"

type UserUseCase interface {
	GetAllUsers() ([]*model.User, error)
	GetOneUser(id int64) (*model.User, error)
	CreateUser(user *model.User) (int64, error)
}

type UserPort interface {
	GetAllUsersOut() ([]*model.User, error)
	GetOneUserOut(ID int64) (*model.User, error)
	CreateUserOut(user *model.User) (int64, error)
}
