package usecases

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/0xfbravo/brla/model"
	"io"
	"net/http"

	"go.uber.org/zap"
)

// Login logs in a super user into BRLA
// See more: https://brla-superuser-api.readme.io/reference/superuserlogin
func (u *Impl) Login(email string, password string) (*model.Session, error) {
	u.log.Info("Logging into BRLA", zap.String("email", email))

	// Parse
	body := map[string]string{
		"email":    email,
		"password": password,
	}
	bodyBytes, _ := json.Marshal(body)

	url := u.baseUrl + "/v1/superuser/login"
	newRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		u.log.Error("Failed to login on BRLA", zap.Error(err))
		return nil, err
	}

	// Send
	client := &http.Client{}
	resp, err := client.Do(newRequest)
	if err != nil {
		u.log.Error("Failed to login on BRLA", zap.Error(err))
		return nil, err
	}

	// Parse
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		u.log.Error("Failed to read response", zap.Error(err))
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		u.log.Error("Failed to login on BRLA", zap.Int("status", resp.StatusCode), zap.String("response", string(respBody)))
		return nil, errors.New("failed to login on BRLA")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			u.log.Error("Failed to close response body", zap.Error(err))
		}
	}(resp.Body)

	var respMap model.Login
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		u.log.Error("Failed to parse response", zap.Error(err), zap.String("response", string(respBody)))
		return nil, err
	}

	// Save session
	newSession := model.Session{
		Email:       email,
		Password:    password,
		AccessToken: respMap.AccessToken,
	}
	err = u.repo.SaveSession(&newSession)
	if err != nil {
		u.log.Warn("Failed to save session", zap.Error(err))
	}

	u.log.Info("BRLA login successful", zap.String("response", string(respBody)))
	return &newSession, nil
}
