package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type HealthCheckRepositoryImpl struct {
	DB *gorm.DB
}

var HealthCheckRepositorySet = wire.NewSet(
	wire.Struct(new(HealthCheckRepositoryImpl), "*"),
	wire.Bind(new(repository.HealthCheckRepository), new(*HealthCheckRepositoryImpl)),
)

func (r *HealthCheckRepositoryImpl) CheckDBAlive() error {
	// database/sqlのPingを使って死活監視
	// https://qiita.com/hgsgtk/items/770c51559f374b36da3f
	return r.DB.DB().Ping()
}
