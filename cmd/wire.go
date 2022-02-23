//+build wireinject

package main

import (
	"github.com/google/wire"
	"my-clean-architecture/api"
	"my-clean-architecture/api/handlers"
	"my-clean-architecture/app"
	"my-clean-architecture/repo"
	"my-clean-architecture/service"
)

func InitServer() *app.Server {
	wire.Build(
		app.InitDB,
		repo.NewMysqlArticleRepository,
		service.NewArticleService,
		handlers.NewArticleHandler,
		api.NewRouter,
		app.NewServer,
		app.NewGinEngine)
	return &app.Server{}
}
