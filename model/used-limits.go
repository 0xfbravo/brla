package model

type UsedLimits struct {
	UsedLimit Limits `json:"usedLimit" validate:"required"`
}
