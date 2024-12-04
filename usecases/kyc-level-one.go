package usecases

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/0xfbravo/brla/enum"
	model2 "github.com/0xfbravo/brla/model"
	"io"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"
)

// KycLevelOne performs KYC Level One on BRLA for both Natural and Legal Person
// See more:
// - https://brla-superuser-api.readme.io/reference/superuserkycpflevel1
// - https://brla-superuser-api.readme.io/reference/superuserkycpjlevel1
// - https://brla-superuser-api.readme.io/reference/passkyclevel1 (Sandbox only)
// - https://brla-superuser-api.readme.io/reference/passkyclevel2 (Sandbox only)
func (u *Impl) KycLevelOne(options *model2.KycLevelOneOptions) (*model2.KycLevelOne, error) {
	u.log.Info("Performing KYC Level One on BRLA", zap.Any("options", options))

	// Check session
	session, err := u.CheckSession()
	if err != nil {
		return nil, err
	}

	parsedDate, err := time.Parse("2006-01-02", options.Birthdate)
	if err != nil {
		u.log.Error("Failed to parse birth date", zap.Error(err))
		return nil, err
	}
	formattedBirthdate := parsedDate.Format("2006-Jan-02")

	// Parse
	var body map[string]string
	if options.PersonType == enum.Individual {
		body = map[string]string{
			"cpf":       options.Document,
			"birthDate": formattedBirthdate,
			"fullName":  options.Name,
		}
	} else {
		body = map[string]string{
			"cnpj":        options.Document,
			"companyName": options.Name,
			"startDate":   formattedBirthdate,
		}
	}
	bodyBytes, _ := json.Marshal(body)

	// If we're not in production, KYC is bypassed in a sandbox endpoint
	var kycEndpoint string
	if u.isProduction {
		kycEndpoint = "/v1/superuser/kyc/" + string(options.PersonType) + "/level1"
	} else {
		kycEndpoint = "/v1/superuser/kyc/" + string(options.PersonType) + "-free/pass-kyc-level1"
	}

	url := u.baseUrl + kycEndpoint
	newRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		u.log.Error("Failed to create KYC Level One request", zap.Error(err))
		return nil, err
	}

	// Send
	newRequest.Header.Set("Authorization", "Bearer "+session.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(newRequest)
	if err != nil {
		u.log.Error("Failed to do KYC Level One on BRLA", zap.Error(err))
		return nil, err
	}

	// Parse
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		u.log.Error("Failed to read response", zap.Error(err))
		return nil, err
	}

	if resp.StatusCode == http.StatusBadRequest && strings.Contains(string(respBody), "user already passed kyc") {
		u.log.Warn("User already passed kyc", zap.String("taxId", options.Document))
		return nil, errors.New("user already passed kyc")
	}

	if resp.StatusCode != http.StatusCreated {
		u.log.Error("Failed to KYC on BRLA", zap.Int("status", resp.StatusCode), zap.Any("body", string(respBody)))
		return nil, errors.New("failed to KYC on BRLA, status code: " + fmt.Sprint(resp.StatusCode))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			u.log.Error("Failed to close response body", zap.Error(err))
		}
	}(resp.Body)

	u.log.Info("Handling KYC Level One response", zap.String("response", string(respBody)))
	var respMap model2.KycLevelOne
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		u.log.Error("Failed to parse response", zap.Error(err))
		return nil, err
	}

	u.log.Info("BRLA KYC Level One performed successfully", zap.Any("response", respMap))
	return &respMap, nil
}
