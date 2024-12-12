package usecases

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

// checkTokenExpiration checks if the token has expired
func (u *Impl) checkTokenExpiration(jwtString string) error {
	u.log.Info("Checking token expiration time")
	token, _, err := jwt.NewParser().ParseUnverified(jwtString, jwt.MapClaims{})
	if err != nil {
		u.log.Error("Failed to parse token", zap.Error(err))
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		expirationTime, err := claims.GetExpirationTime()
		if err != nil {
			u.log.Error("Failed to get expiration time", zap.Error(err))
			return err
		}

		if time.Now().After(expirationTime.Time) {
			u.log.Warn("Token has expired")
			return errors.New("token has expired")
		}
	}

	u.log.Info("Token has a valid expiration time")
	return nil
}
