package in

import "rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"

type TrackUseCase interface {
	GetAll() ([]model.Music, error)
	GetOne(ID uint64) (model.Music, error)
	Create(model.Music) error
	UpdateOne(ID uint64, m model.Music) error
	SoftDeleteOne(ID uint64) error
}

type TrackPort interface {
	GetAllOut() ([]model.Music, error)
	GetOneOut(ID uint64) (model.Music, error)
	CreateOut(model.Music) error
	UpdateOneOut(ID uint64, m model.Music) error
	SoftDeleteOneOut(ID uint64) error
}
