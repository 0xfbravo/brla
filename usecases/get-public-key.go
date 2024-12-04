package usecases

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"github.com/0xfbravo/brla/model"
	"io"
	"net/http"

	"go.uber.org/zap"
)

// GetPublicKey gets the public key from BRLA
// See more:
// - https://brla-superuser-api.readme.io/reference/pubkey
func (u *Impl) GetPublicKey() (*ecdsa.PublicKey, error) {
	u.log.Info("Getting public key from BRLA")

	req, _ := http.NewRequest("GET", u.baseUrl+"/v1/pubkey", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		u.log.Error("Failed to get public key", zap.Error(err))
		return nil, err
	}

	pubKeyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		u.log.Error("Failed to read response", zap.Error(err))
		return nil, err
	}

	var pubKeyData model.PublicKey
	err = json.Unmarshal(pubKeyBytes, &pubKeyData)
	if err != nil {
		u.log.Error("Failed to unmarshal public key", zap.Error(err))
		return nil, err
	}

	pubBlock, _ := pem.Decode([]byte(pubKeyData.PublicKey))
	if pubBlock == nil {
		u.log.Error("Failed to decode public key")
		return nil, err
	}

	pubKeyUntyped, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		u.log.Error("Failed to parse public key", zap.Error(err))
		return nil, err
	}

	pubKey, ok := pubKeyUntyped.(*ecdsa.PublicKey)
	if !ok {
		u.log.Error("Failed to cast public key")
		return nil, err
	}

	u.log.Info("Public key retrieved successfully", zap.Any("publicKey", pubKeyData.PublicKey))
	return pubKey, nil
}
