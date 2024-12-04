package model

import (
	"github.com/0xfbravo/brla/enum"
	"time"
)

type StaticPix struct {
	Amount                  int64      `json:"amount" validate:"required" example:"1000"`
	WalletAddress           string     `json:"walletAddress" validate:"required" example:"0x0471d29143A6af8DcEf2263573Cb44C3d998764F"`
	Chain                   enum.Chain `json:"chain" validate:"required" example:"Polygon"`
	Due                     *time.Time `json:"due,omitempty" validate:"omitempty" example:"2022-12-31T23:59:59Z"`
	CoverFeeWithBrlaAccount *bool      `json:"coverFeeWithBrlaAccount,omitempty" validate:"omitempty" example:"true"`
	ExternalId              *string    `json:"externalId,omitempty" validate:"omitempty" example:"123456"`
}
