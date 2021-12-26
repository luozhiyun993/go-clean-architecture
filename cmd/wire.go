//+build wireinject

package main

import (
	"github.com/google/wire"
	"my-clean-rchitecture/api"
	"my-clean-rchitecture/api/handlers"
	"my-clean-rchitecture/app"
	"my-clean-rchitecture/repo"
	"my-clean-rchitecture/service"
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
