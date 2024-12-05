package model

import "github.com/0xfbravo/brla/enum"

type WebhookResponse struct {
	Subscription enum.Subscription      `json:"subscription"`
	CreatedAt    int64                  `json:"createdAt"`
	ID           string                 `json:"id"`
	Data         map[string]interface{} `json:"data"`
}
