package repository

import (
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"golang.org/x/net/context"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"
	port "rpolnx.com.br/crud-sql/internal/crud-sql/domain/ports"
	"rpolnx.com.br/crud-sql/internal/crud-sql/infrastructure/repository/postgres"
)

type userRepository struct {
	Db *bun.DB
}

func (svc *userRepository) GetAllUsersOut() ([]*model.User, error) {
	users := make([]*model.User, 0)

	res, err := svc.Db.NewSelect().
		Model(&users).
		Exec(context.Background())

	logrus.Debugf("Result from select %v", res)

	return users, postgres.HandleDbError(err)
}

func (svc *userRepository) GetOneUserOut(id *int64) (*model.User, error) {
	user := &model.User{
		Id: id,
	}

	res, err := svc.Db.NewSelect().
		Model(user).
		WherePK().
		Exec(context.Background())

	logrus.Debugf("Result from select %v", res)

	return user, postgres.HandleDbError(err)
}

func (svc *userRepository) CreateUserOut(user *model.User) (*int64, error) {
	user.Id = nil

	insert, err := svc.Db.NewInsert().
		Model(user).
		Exec(context.Background())

	if err != nil {
		return nil, postgres.HandleDbError(err)
	}

	affected, err := insert.RowsAffected()
	logrus.Info(affected)

	return user.Id, err
}

func (svc *userRepository) UpdateUserOut(id *int64, user *model.User) (*int64, error) {
	user.Id = id

	updated, err := svc.Db.NewUpdate().
		Model(user).
		WherePK().
		OmitZero().
		Exec(context.Background())

	if err != nil {
		return nil, postgres.HandleDbError(err)
	}

	affected, err := updated.RowsAffected()

	logrus.Infof("Afftected %d", affected)

	if affected == 0 {
		return nil, &config.NotFoundError{Name: "User"}
	}

	return user.Id, err
}

func (svc *userRepository) DeleteUserOut(id *int64) error {
	user := &model.User{Id: id}

	deleted, err := svc.Db.NewDelete().
		Model(user).
		WherePK().
		//OmitZero().
		Exec(context.Background())

	if err != nil {
		return err
	}

	affected, err := deleted.RowsAffected()
	logrus.Info(affected)

	if affected == 0 {
		return &config.NotFoundError{Name: "User"}
	}

	return postgres.HandleDbError(err)
}

func NewUserRepository(db *bun.DB) port.UserPort {
	return &userRepository{
		Db: db,
	}
}
