package postgres

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"rpolnx.com.br/crud-sql/internal/crud-sql/application/config"
	"rpolnx.com.br/crud-sql/internal/crud-sql/domain/model"
	"rpolnx.com.br/crud-sql/internal/crud-sql/infrastructure/repository"
	"time"
)

type pgRepository struct {
	client   *gorm.DB
	database string
	timeout  time.Duration
}

func (pg *pgRepository) FindAllTracks() ([]model.Music, error) {
	return nil, nil
}
func (pg *pgRepository) GetOneTrack(ID uint64) (*model.Music, error) {
	return nil, nil
}
func (pg *pgRepository) CreateTrack(m *model.Music) error {
	return nil
}
func (pg *pgRepository) UpdateOneTrack(ID uint64, m *model.Music) error {
	return nil
}
func (pg *pgRepository) DeleteOneTrack(ID uint64) error {
	return nil
}

func InitializeTrackRepo(cfg *config.Configuration) (repository.TrackRepository, error) {
	initialDbTime := time.Now()

	logrus.Infof("[PG DB] Initializing store dependencies")
	repo, err := newTrackRepo(cfg.Db)

	if err != nil {
		logrus.Errorf("[PG DB] Error initializing store dependencies")
		return nil, err
	}

	deltaMongo := time.Since(initialDbTime).Milliseconds()
	logrus.Infof("[PG DB] Finalized store dependency in %dms", deltaMongo)

	return repo, nil
}

func newTrackRepo(dbCfg config.Db) (repository.TrackRepository, error) {
	repo := &pgRepository{
		timeout:  time.Duration(dbCfg.Timeout) * time.Second,
		database: dbCfg.DbName,
	}

	client, err := NewPgClient(dbCfg)
	if err != nil {
		return nil, err
	}

	repo.client = client

	return repo, nil
}
