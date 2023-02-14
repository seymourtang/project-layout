package injector

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/redis/go-redis/v9"

	"github.com/seymourtang/project-layout/internal/repository"
	"github.com/seymourtang/project-layout/internal/task"
)

type injector struct {
	studentRepository repository.Student
	redisClient       redis.UniversalClient
	taskGroup         []task.Runner
}

func NewInjector(
	studentRepository repository.Student,
	redisClient redis.UniversalClient,
	taskGroup []task.Runner,
) *injector {
	return &injector{
		studentRepository: studentRepository,
		redisClient:       redisClient,
		taskGroup:         taskGroup,
	}
}

func (i *injector) Run() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for _, runner := range i.taskGroup {
		log.Printf("[%s] is started", runner.Name())
		if err := runner.Start(context.Background()); err != nil {
			panic(err)
		}
	}
	<-c
	for _, runner := range i.taskGroup {
		if err := runner.Stop(context.Background()); err != nil {
		} else {
			log.Printf("[%s] is stopped", runner.Name())
		}
	}
}
