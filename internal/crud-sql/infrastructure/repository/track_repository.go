package repository

import "rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"

type TrackRepository interface {
	FindAllTracks() ([]model.Music, error)
	GetOneTrack(ID uint64) (*model.Music, error)
	CreateTrack(m *model.Music) error
	UpdateOneTrack(ID uint64, m *model.Music) error
	DeleteOneTrack(ID uint64) error
}
