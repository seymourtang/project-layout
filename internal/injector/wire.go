//go:generate wire
//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"

	"github.com/seymourtang/project-layout/cmd/app/option"
	"github.com/seymourtang/project-layout/internal/data/cache"
	"github.com/seymourtang/project-layout/internal/data/db"
	"github.com/seymourtang/project-layout/internal/mq"
	"github.com/seymourtang/project-layout/internal/repository"
	"github.com/seymourtang/project-layout/internal/server"
	"github.com/seymourtang/project-layout/internal/task"
)

func Build() (*injector, func(), error) {
	panic(wire.Build(
		repository.ProviderSet,
		db.ProviderSet,
		option.ProviderSet,
		server.ProviderSet,
		cache.ProviderSet,
		task.ProviderSet,
		mq.ProviderSet,
		NewInjector,
	))
}
