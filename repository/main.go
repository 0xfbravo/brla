package repository

import (
	"github.com/0xfbravo/brla/interfaces"
	"go.uber.org/zap"
)

type repositoryImpl struct {
	email       string
	password    string
	accessToken *string
	log         *zap.Logger
}

func New(
	log *zap.Logger,
) interfaces.Repository {
	repo := repositoryImpl{
		log: log,
	}

	return &repo
}
