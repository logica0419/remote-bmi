//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"github.com/logica0419/remote-bmi/server/repository"
	"github.com/logica0419/remote-bmi/server/router"
)

var set = wire.NewSet(
	newRepositoryConfig,
	repository.NewRepository,
	repository.GetSqlDB,

	newRouterConfig,
	router.NewRouter,
)

func setupRouter(cfg *Config) (*router.Router, error) {
	wire.Build(set)
	return nil, nil
}
