package usecases

import (
	"errors"
	"github.com/0xfbravo/brla/abi"
	"github.com/0xfbravo/brla/model"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"go.uber.org/zap"
)

// GetBalanceOf retrieves the balance of a wallet directly from the blockchain
func (u *Impl) GetBalanceOf(options *model.BalanceOfOptions) (*model.BalanceOf, error) {
	u.log.Info("Retrieving balanceOf(wallet)", zap.Any("options", options))

	tokenAddresses, err := u.GetContractAddress(options.Chain)
	if err != nil {
		u.log.Error("Failed to get contract address", zap.Error(err))
		return nil, err
	}

	client, err := ethclient.Dial(options.RpcUrl)
	if err != nil {
		u.log.Error("Failed to connect to the Ethereum client", zap.Error(err))
		return nil, errors.New("failed to connect to the Ethereum client")
	}

	contractAddress := common.HexToAddress(tokenAddresses.BRLATokenAddress)
	contract, err := brla.NewBrla(contractAddress, client)
	if err != nil {
		u.log.Error("Failed to get contract", zap.Error(err))
		return nil, err
	}

	balance, err := contract.BalanceOf(nil, options.Wallet)
	if err != nil {
		u.log.Error("Failed to get balance", zap.Error(err), zap.Any("options", options))
		return nil, err
	}

	// Convert balance to cents
	newWei := new(big.Int).Set(balance)
	newWei.Mul(newWei, new(big.Int).SetInt64(100))

	// Convert balance to ether
	ether := new(big.Float).SetInt(newWei)       // Convert big.Int (wei) to big.Float
	ether.Quo(ether, big.NewFloat(params.Ether)) // Divide by 10^18 (Ether constant)
	etherInt, _ := ether.Int64()                 // Convert big.Float to big.Int

	u.log.Info("Wallet balance retrieved successfully", zap.Any("balance", balance), zap.Any("ether", etherInt), zap.Any("options", options))
	return &model.BalanceOf{
		Balance:    etherInt,
		BalanceWei: balance.String(),
	}, nil
}
