package interfaces

import (
	"github.com/0xfbravo/brla/model"
)

type Repository interface {
	// GetSession gets the session token from the previous login if it's available
	GetSession() (*model.Session, error)

	// SaveSession saves the session from the provided email, password and last access token
	SaveSession(session *model.Session) error
}
