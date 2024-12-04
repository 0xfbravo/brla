package model

type BalanceOf struct {
	Balance    int64  `json:"balance,omitempty" validate:"required" example:"100000000"`
	BalanceWei string `json:"balanceWei,omitempty" validate:"required" example:"1000000000000000000"`
}
