package server

import (
	"github.com/google/wire"

	"github.com/seymourtang/project-layout/internal/server/http"
)

var ProviderSet = wire.NewSet(http.ProviderSet)
