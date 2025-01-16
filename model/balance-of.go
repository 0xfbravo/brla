package model

type BalanceOf struct {
	Balance    int64  `json:"balance" validate:"required" example:"100000000"`
	BalanceWei string `json:"balanceWei" validate:"required" example:"1000000000000000000"`
}
