package model

import "time"

type KYC struct {
	ID           string    `json:"id" validate:"required" example:"00000000-0000-0000-0000-000000000000"`
	Level        int       `json:"level" validate:"required" example:"1"`
	DocumentData string    `json:"documentData" validate:"required" example:"000.000.000-00"`
	DocumentType string    `json:"documentType" validate:"required" example:"CPF"`
	Limits       Limits    `json:"limits" validate:"required"`
	CreatedAt    time.Time `json:"createdAt" validate:"required" example:"2000-01-01T00:00:00Z"`
}
