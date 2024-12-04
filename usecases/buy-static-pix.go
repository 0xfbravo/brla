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

// BuyStaticPix creates a new static pix on BRLA
// See more:
// - https://brla-superuser-api.readme.io/reference/superuserbuystaticpix
func (u *Impl) BuyStaticPix(options *model.BuyStaticPixOptions) (*model.Pix, error) {
	u.log.Info("Buying static pix on BRLA", zap.Any("options", options))

	// Check session
	session, err := u.CheckSession()
	if err != nil {
		return nil, err
	}

	// Parse
	body := model.StaticPix{
		Amount:                  options.Amount,
		WalletAddress:           options.WalletAddress,
		Chain:                   options.Chain,
		Due:                     options.Due,
		CoverFeeWithBrlaAccount: options.CoverFeeWithBrlaAccount,
		ExternalId:              options.ExternalId,
	}
	bodyBytes, _ := json.Marshal(body)

	url := u.baseUrl + "/v1/superuser/buy/static-pix?taxId=" + options.DocumentNumber
	newRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		u.log.Error("Failed to buy static pix on BRLA", zap.Error(err))
		return nil, err
	}

	// Send
	newRequest.Header.Set("Authorization", "Bearer "+session.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(newRequest)
	if err != nil {
		u.log.Error("Failed to buy static pix on BRLA", zap.Error(err))
		return nil, err
	}

	// Parse
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		u.log.Error("Failed to read response", zap.Error(err))
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		u.log.Error("Failed to buy static pix on BRLA", zap.Any("status", resp.StatusCode), zap.String("response", string(respBody)))
		return nil, errors.New("failed to buy static pix on BRLA")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			u.log.Error("Failed to close response body", zap.Error(err))
		}
	}(resp.Body)

	var respMap model.Pix
	if err := json.Unmarshal(respBody, &respMap); err != nil {
		u.log.Error("Failed to parse response", zap.Error(err), zap.String("response", string(respBody)))
		return nil, err
	}

	u.log.Info("BRLA static pix bought successfully", zap.Any("options", options), zap.Any("response", respMap))
	return &respMap, nil
}
