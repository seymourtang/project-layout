package cache

import (
	"github.com/google/wire"

	"github.com/seymourtang/project-layout/internal/data/cache/redis"
)

var ProvideSet = wire.NewSet(redis.New)
