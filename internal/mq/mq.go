package mq

import (
	"github.com/google/wire"

	"github.com/seymourtang/project-layout/internal/mq/redis"
)

var ProviderSet = wire.NewSet(redis.NewChannelMQ)
