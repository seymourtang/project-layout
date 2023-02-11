package task

import (
	"context"

	"github.com/google/wire"

	"github.com/seymourtang/project-layout/internal/mq/redis"
	"github.com/seymourtang/project-layout/internal/server/http"
)

var ProviderSet = wire.NewSet(NewGroup)

func NewGroup(server *http.Server, channel *redis.ChannelMQ) []Runner {
	return []Runner{
		server,
		channel,
	}
}

type Runner interface {
	Name() string
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
