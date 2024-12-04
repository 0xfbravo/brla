package model

import (
	"github.com/0xfbravo/brla/enum"
)

type KycLevelTwo struct {
	ID           string            `json:"id,omitempty" example:"3f056379-7cc8-45c4-a987-ed7d90aade8b"`
	DocumentType enum.DocumentType `json:"documentType,omitempty" validate:"required,oneof=RG CNH" enum:"RG,CNH" example:"RG"`
	Selfie       string            `json:"selfieBase64,omitempty" example:"base64_encoded_selfie"`
	RGFront      string            `json:"rgFrontBase64,omitempty" example:"base64_encoded_rgfront"`
	RGBack       string            `json:"rgBackBase64,omitempty" example:"base64_encoded_rgback"`
	CNH          string            `json:"cnhBase64,omitempty" example:"base64_encoded_cnh"`
}
