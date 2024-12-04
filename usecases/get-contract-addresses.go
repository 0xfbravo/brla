package usecases

import (
	"encoding/json"
	"errors"
	"github.com/0xfbravo/brla/enum"
	"github.com/0xfbravo/brla/model"
	"io"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

// GetContractAddress gets the contract addresses from BRLA for the given chain
// See more:
// - https://brla-superuser-api.readme.io/reference/smartcontractaddresses
func (u *Impl) GetContractAddress(chain enum.Chain) (*model.TokenAddresses, error) {
	u.log.Info("Getting contract addresses from BRLA", zap.Any("chain", chain))

	req, _ := http.NewRequest("GET", u.baseUrl+"/v1/addresses", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		u.log.Error("Failed to get public key", zap.Error(err))
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		u.log.Error("Failed to read response", zap.Error(err))
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		u.log.Error("Failed to get contract addresses", zap.Any("status", resp.StatusCode), zap.String("response", string(respBody)))
		return nil, errors.New("failed to get contract addresses")
	}

	var contractAddresses map[string]model.TokenAddresses
	err = json.Unmarshal(respBody, &contractAddresses)
	if err != nil {
		u.log.Error("Failed to unmarshal contract addresses", zap.Error(err))
		return nil, err
	}

	contractAddress, ok := contractAddresses[strings.ToLower(string(chain))]
	if !ok {
		u.log.Error("Failed to get contract address", zap.Any("chain", chain))
		return nil, errors.New("failed to get contract address")
	}

	u.log.Info("Contract addresses retrieved successfully", zap.Any("contractAddress", contractAddress), zap.Any("chain", chain))
	return &contractAddress, nil
}
