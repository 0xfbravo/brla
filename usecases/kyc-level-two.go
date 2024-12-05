package usecases

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/0xfbravo/brla/enum"
	"github.com/0xfbravo/brla/model"
	"io"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

// KycLevelTwo performs KYC Level Two on BRLA for both Natural and Legal Person
// See more:
// - https://brla-superuser-api.readme.io/reference/superuserkycpflevel2
// - https://brla-superuser-api.readme.io/reference/superuserkycpjlevel2
func (u *Impl) KycLevelTwo(options *model.KycLevelTwoOptions) error {
	u.log.Info("Performing KYC Level Two on BRLA", zap.Any("options", options))

	// Check session
	session, err := u.CheckSession()
	if err != nil {
		return err
	}

	if (options.PersonType == enum.LegalPerson && options.PartnerCpf == nil) || (options.PersonType == enum.LegalPerson && options.PartnerFullName == nil) {
		u.log.Error("Individual partner CPF and full name are required for Legal Person KYC Level Two", zap.Any("options", options))
		return errors.New("individual partner CPF and full name are required for Legal Person KYC Level Two")
	}

	// Parse
	var body map[string]string
	if options.PersonType == enum.Individual {
		body = map[string]string{
			"documentType": string(options.Files.DocumentType),
			"cpf":          options.Document,
		}
	} else {
		body = map[string]string{
			"documentType": string(options.Files.DocumentType),
			"cpf":          *options.PartnerCpf,
			"cnpj":         options.Document,
			"fullName":     *options.PartnerFullName,
		}
	}
	bodyBytes, _ := json.Marshal(body)

	// If we're not in production, KYC is bypassed in a sandbox endpoint
	var kycEndpoint string
	if u.isProduction {
		kycEndpoint = "/v1/superuser/kyc/" + string(options.PersonType) + "/level2"
	} else {
		kycEndpoint = "/v1/superuser/kyc/" + string(options.PersonType) + "-free/pass-kyc-level2"
	}

	url := u.baseUrl + kycEndpoint
	newRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		u.log.Error("Failed to create KYC Level Two request", zap.Error(err))
		return err
	}

	// Send
	newRequest.Header.Set("Authorization", "Bearer "+session.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(newRequest)
	if err != nil {
		u.log.Error("Failed to do KYC Level Two on BRLA", zap.Error(err))
		return err
	}

	// Parse
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		u.log.Error("Failed to read response", zap.Error(err))
		return err
	}

	if resp.StatusCode == http.StatusBadRequest && strings.Contains(string(respBody), "user already passed kyc") {
		u.log.Warn("User already passed kyc", zap.Any("options", options), zap.Any("response", string(respBody)))
		return errors.New("user already passed kyc")
	}

	if resp.StatusCode != http.StatusCreated {
		u.log.Error("Failed to KYC on BRLA", zap.Int("status", resp.StatusCode), zap.Any("body", string(respBody)))
		return errors.New("failed to KYC on BRLA, status code: " + fmt.Sprint(resp.StatusCode))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			u.log.Error("Failed to close response body", zap.Error(err))
		}
	}(resp.Body)

	u.log.Info("Handling KYC Level Two response", zap.String("response", string(respBody)))
	var respMap map[string]string
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		u.log.Error("Failed to parse response", zap.Error(err))
		return err
	}

	if u.isProduction {
		u.log.Warn("Bypassing KYC file upload in sandbox environment", zap.Any("response", respMap))
		return nil
	}

	err = u.uploadImage(respMap["selfieUploadUrl"], options.Files.Selfie)
	if err != nil {
		u.log.Error("Failed to upload images selfie", zap.Error(err))
		return err
	}

	if options.Files.DocumentType == enum.RG {
		err = u.uploadImage(respMap["RGFrontUploadUrl"], options.Files.RGBack)
		if err != nil {
			u.log.Error("Failed to upload images rgback", zap.Error(err))
			return err
		}

		err = u.uploadImage(respMap["RGBackUploadUrl"], options.Files.RGFront)
		if err != nil {
			u.log.Error("Failed to upload images rgfront", zap.Error(err))
			return err
		}
	} else {
		err = u.uploadImage(respMap["CNHUploadUrl"], options.Files.CNH)
		if err != nil {
			u.log.Error("Failed to upload images cnh", zap.Error(err))
			return err
		}
	}

	u.log.Info("BRLA KYC Level Two performed successfully", zap.Any("response", respMap))
	return nil
}
