package model

type Limits struct {
	LimitMint              int64 `json:"limitMint" validate:"required" example:"100000"`
	LimitBurn              int64 `json:"limitBurn" validate:"required" example:"100000"`
	LimitSwapBuy           int64 `json:"limitSwapBuy" validate:"required" example:"100000"`
	LimitSwapSell          int64 `json:"limitSwapSell" validate:"required" example:"100000"`
	LimitBRLAOutOwnAccount int64 `json:"limitBRLAOutOwnAccount" validate:"required" example:"100000"`
	LimitBRLAOutThirdParty int64 `json:"limitBRLAOutThirdParty" validate:"required" example:"100000"`
}
