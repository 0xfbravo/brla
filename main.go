package brla

import (
	"github.com/0xfbravo/brla/interfaces"
	"github.com/0xfbravo/brla/repository"
	"github.com/0xfbravo/brla/usecases"
	"go.uber.org/zap"
)

type Module struct {
	repository interfaces.Repository
	UseCases   interfaces.UseCases
}

// New creates a new module instance
func New(
	isProduction bool,
	log *zap.Logger,
) Module {
	var baseUrl string
	if isProduction {
		baseUrl = "https://api.brla.digital:5567"
	} else {
		baseUrl = "https://api.brla.digital:4567"
	}

	r := repository.New(log)
	u := usecases.New(isProduction, baseUrl, r, log)
	return Module{
		repository: r,
		UseCases:   u,
	}
}
