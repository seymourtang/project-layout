//go:generate wire
//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"

	"github.com/seymourtang/project-layout/cmd/app/option"
	"github.com/seymourtang/project-layout/internal/data/cache"
	"github.com/seymourtang/project-layout/internal/data/db"
	"github.com/seymourtang/project-layout/internal/repository"
	"github.com/seymourtang/project-layout/internal/server"
)

func Build() (*Injector, func(), error) {
	panic(wire.Build(
		repository.ProvideSet,
		db.ProvideSet,
		option.ProvideSet,
		server.ProviderSet,
		cache.ProvideSet,
		NewInjector,
	))
}
