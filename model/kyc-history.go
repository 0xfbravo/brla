package model

type KYCHistory struct {
	History []KYC `json:"kycs" validate:"required"`
}
