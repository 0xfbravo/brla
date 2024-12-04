package model

import (
	"github.com/0xfbravo/brla/enum"
	"github.com/ethereum/go-ethereum/common"
)

type BalanceOfOptions struct {
	RpcUrl string
	Chain  enum.Chain
	Wallet common.Address
}
