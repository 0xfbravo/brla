package model

type WebhookResponse struct {
	Subscription string                 `json:"subscription"`
	CreatedAt    int64                  `json:"createdAt"`
	ID           string                 `json:"id"`
	Data         map[string]interface{} `json:"data"`
}
