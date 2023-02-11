package service

import "context"

type Student interface {
	Get(ctx context.Context, ID string)
}
