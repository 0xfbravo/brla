package model

import "github.com/0xfbravo/brla/enum"

type SellOptions struct {
	TaxId                   string
	PixKey                  string
	WalletAddress           string
	Chain                   enum.Chain
	Amount                  int
	ReferenceLabel          string
	ExternalId              string
	CoverFeeWithBrlaAccount bool
	Signature               string
	SignatureDeadline       int
}
