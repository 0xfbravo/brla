package usecases

import (
	"errors"
	"github.com/0xfbravo/brla/model"
)

// CheckSession checks if the client is authorized to access BRLA
func (u *Impl) CheckSession() (*model.Session, error) {
	session, err := u.repo.GetSession()
	if err != nil {
		u.log.Error("Unauthorized access to BRLA, access token is empty. Please login first")
		return nil, errors.New("unauthorized access to BRLA")
	}

	u.log.Info("Checking if current access token is still valid")
	err = u.checkTokenExpiration(session.AccessToken)
	if err == nil {
		u.log.Info("Client still authorized on BRLA")
		return session, nil
	}

	u.log.Warn("Access token expired, re-authorizing on BRLA")
	newSession, err := u.Login(session.Email, session.Password)
	if err != nil {
		u.log.Error("Failed to re-authorize client on BRLA")
		return nil, err
	}

	u.log.Info("Client re-authorized on BRLA successfully")
	return newSession, nil
}
