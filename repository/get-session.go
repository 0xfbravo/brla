package repository

import (
	"errors"
	"github.com/0xfbravo/brla/model"
)

// GetSession gets the session token from the previous login if it's available
func (r *repositoryImpl) GetSession() (*model.Session, error) {
	r.log.Info("Getting session")

	if r.email == "" {
		r.log.Error("Email not found")
		return nil, errors.New("email not found")
	}

	if r.password == "" {
		r.log.Error("Password not found")
		return nil, errors.New("password not found")
	}

	if r.accessToken == nil || *r.accessToken == "" {
		r.log.Error("Generated access token not found")
		return nil, errors.New("generated access token not found")
	}

	r.log.Info("Session retrieved successfully")
	return &model.Session{
		Email:       r.email,
		Password:    r.password,
		AccessToken: *r.accessToken,
	}, nil
}
