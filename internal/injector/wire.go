//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"

	"github.com/seymourtang/project-layout/internal/data/cache"
	"github.com/seymourtang/project-layout/internal/server"
)

func Build() (*Injector, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		cache.ProvideSet,
		NewInjector,
	))
}
