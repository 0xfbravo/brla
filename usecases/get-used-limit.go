package usecases

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/0xfbravo/brla/model"
	"io"

	"net/http"

	"go.uber.org/zap"
)

// GetUsedLimit retrieves user used limit from BRLA
// See more:
// - https://brla-superuser-api.readme.io/reference/superuserusedlimit
func (u *Impl) GetUsedLimit(taxId string) (*model.Limits, error) {
	u.log.Info("Retrieving user used limit from BRLA", zap.String("taxId", taxId))

	// Check session
	session, err := u.CheckSession()
	if err != nil {
		return nil, err
	}

	url := u.baseUrl + "/v1/superuser/used-limit?taxId=" + taxId
	newRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		u.log.Error("Failed to get user used limit on BRLA", zap.Error(err), zap.String("taxId", taxId))
		return nil, err
	}

	// Send
	newRequest.Header.Set("Authorization", "Bearer "+session.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(newRequest)
	if err != nil {
		u.log.Error("Failed to get user used limit on BRLA", zap.Error(err), zap.String("taxId", taxId))
		return nil, err
	}

	// Parse
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		u.log.Error("Failed to read response", zap.Error(err))
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		u.log.Error("Failed to get user used limit on BRLA", zap.Int("status", resp.StatusCode), zap.String("response", string(respBody)))
		return nil, errors.New("failed to get user used limit on BRLA, status code: " + fmt.Sprint(resp.StatusCode))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			u.log.Error("Failed to close response body", zap.Error(err))
		}
	}(resp.Body)

	var respMap model.UsedLimits
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		u.log.Error("Failed to parse response", zap.Error(err), zap.String("response", string(respBody)))
		return nil, err
	}

	u.log.Info("BRLA user used limit retrieved successfully", zap.String("taxId", taxId))
	return &respMap.UsedLimit, nil
}
