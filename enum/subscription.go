package enum

type Subscription string

const (
	Mint          Subscription = "MINT"
	Burn          Subscription = "BURN"
	MoneyTransfer Subscription = "MONEY-TRANSFER"
	Kyc           Subscription = "KYC"
)
