package usecases

import (
	"github.com/0xfbravo/brla/interfaces"
	"go.uber.org/zap"
)

type Impl struct {
	isProduction bool
	baseUrl      string
	repo         interfaces.Repository
	log          *zap.Logger
}

// New creates a new use cases instance
func New(
	isProduction bool,
	baseUrl string,
	repo interfaces.Repository,
	log *zap.Logger,
) interfaces.UseCases {
	return &Impl{
		isProduction: isProduction,
		baseUrl:      baseUrl,
		repo:         repo,
		log:          log,
	}
}
