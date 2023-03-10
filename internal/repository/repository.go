package repository

import (
	"context"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewStudentImpl,
	wire.Bind(new(Student), new(*studentImpl)),
)

type Student interface {
	Get(ctx context.Context, ID string) (*StudentDTO, error)
}
