package injector

import (
	"context"
	"net/http"

	"github.com/redis/go-redis/v9"

	"github.com/seymourtang/project-layout/internal/repository"
)

type Injector struct {
	studentRepository repository.Student
	redisClient       redis.UniversalClient
	httpServer        *http.ServeMux
}

func NewInjector(
	studentRepository repository.Student,
	redisClient redis.UniversalClient,
	httpServer *http.ServeMux,
) *Injector {
	return &Injector{
		studentRepository: studentRepository,
		redisClient:       redisClient,
		httpServer:        httpServer,
	}
}

func (i *Injector) Run() {
	i.studentRepository.Get(context.Background(), "232")
}
