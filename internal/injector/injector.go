package injector

import (
	"net/http"

	"github.com/redis/go-redis/v9"
)

type Injector struct {
	httpServer  *http.ServeMux
	redisClient redis.UniversalClient
}

func NewInjector(httpServer *http.ServeMux, redisClient redis.UniversalClient) *Injector {
	return &Injector{
		httpServer:  httpServer,
		redisClient: redisClient,
	}
}

func (i *Injector) Run() {
}
