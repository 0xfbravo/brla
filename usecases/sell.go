package usecases

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/0xfbravo/brla/model"
	"go.uber.org/zap"
)

// Sell sells tokens on BRLA
// See more: https://brla-superuser-api.readme.io/reference/superusersell
func (u *Impl) Sell(options *model.SellOptions) (*string, error) {
	// Check session
	session, err := u.CheckSession()
	if err != nil {
		return nil, err
	}

	// Parse
	payload := map[string]interface{}{
		"chain":                   string(options.Chain),
		"pixKey":                  options.PixKey,
		"amount":                  options.Amount,
		"walletAddress":           options.WalletAddress,
		"referenceLabel":          options.ReferenceLabel,
		"coverFeeWithBrlaAccount": options.CoverFeeWithBrlaAccount,
		"signature":               options.Signature,
		"signatureDeadline":       options.SignatureDeadline,
	}
	bodyBytes, _ := json.Marshal(payload)

	url := u.baseUrl + "/superuser/sell?taxId=" + options.TaxId
	newRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		u.log.Error("Failed to create sell request on BRLA", zap.Error(err))
		return nil, err
	}

	// Parse
	newRequest.Header.Set("Authorization", "Bearer "+session.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(newRequest)
	if err != nil {
		u.log.Error("Failed to sell tokens on BRLA", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		u.log.Error("Failed to read response", zap.Error(err))
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		u.log.Error("Failed to sell on BRLA", zap.Int("status", resp.StatusCode), zap.String("body", string(body)))
		return nil, errors.New("failed to sell on BRLA")
	}

	var bodyData model.Sell
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		u.log.Error("Failed to parse response", zap.Error(err))
		return nil, err
	}

	u.log.Info("Sell on BRLA performed successfully", zap.Any("body", bodyData))
	return &bodyData.ID, nil
}
