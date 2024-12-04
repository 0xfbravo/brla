package usecases

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

// uploadImage uploads a KYC image to a given URL
func (u *Impl) uploadImage(url string, data string) error {
	u.log.Info("Uploading KYC image", zap.String("url", url))
	imageData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		u.log.Error("Failed to decode image data", zap.Error(err))
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(imageData))
	if err != nil {
		u.log.Error("Failed to create request", zap.Error(err), zap.String("uploadURL", url))
		return err
	}
	req.Header.Set("Content-Type", "application/octet-stream") // Set appropriate content type for binary data

	u.log.Info("Sending KYC image upload request", zap.String("uploadURL", url))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		u.log.Error("Failed to upload image", zap.Error(err), zap.String("uploadURL", url))
		return err
	}
	u.log.Info("KYC image upload request sent successfully", zap.String("uploadURL", url))

	// Parse
	u.log.Info("Reading response body", zap.String("uploadURL", url))
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		u.log.Error("Failed to read response", zap.Error(err), zap.String("uploadURL", url))
		return nil
	}
	u.log.Info("Response body read successfully", zap.String("response", string(respBody)), zap.String("uploadURL", url))

	if resp.StatusCode != http.StatusOK {
		u.log.Error("Failed to upload KYC image", zap.Int("status", resp.StatusCode), zap.String("response", string(respBody)), zap.String("uploadURL", url))
		return errors.New("failed to upload KYC image, status code: " + fmt.Sprint(resp.StatusCode))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			u.log.Error("Failed to close response body", zap.Error(err))
		}
	}(resp.Body)

	u.log.Info("BRLA KYC image uploaded successfully", zap.String("response", string(respBody)), zap.String("uploadURL", url))
	return nil
}
