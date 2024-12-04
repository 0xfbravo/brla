package repository

import (
	"github.com/0xfbravo/brla/model"
	"go.uber.org/zap"
)

// SaveSession saves the session from the provided email, password and last access token
func (r *repositoryImpl) SaveSession(session *model.Session) error {
	r.log.Info("Saving BRLA session", zap.Any("session", session))
	// TODO(0xfbravo): Save sessions in a secure storage
	r.email = session.Email
	r.password = session.Password
	r.accessToken = &session.AccessToken
	r.log.Info("BRLA session saved successfully", zap.Any("session", session))
	return nil
}
