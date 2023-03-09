package repository

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"
	port "rpolnx.com.br/crud-sql/internal/crud-sql/domain/ports"
)

type userRepository struct {
	Db *pg.DB
}

func (svc *userRepository) GetAllUsersOut() ([]*model.User, error) {
	users := make([]*model.User, 0)

	err := svc.Db.Model(users).
		Select()

	return users, err
}

func (svc *userRepository) GetOneUserOut(id int64) (*model.User, error) {
	user := &model.User{
		Id: id,
	}

	err := svc.Db.Model(user).
		WherePK().
		Select()

	return user, err
}

func (svc *userRepository) CreateUserOut(user *model.User) (int64, error) {

	insert, err := svc.Db.Model(user).
		Insert()

	affected := insert.RowsAffected()
	fmt.Println(affected)

	return user.Id, err
}

func NewUserRepository(db *pg.DB) port.UserPort {
	return &userRepository{
		Db: db,
	}
}
