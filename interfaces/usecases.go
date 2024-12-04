package interfaces

import (
	model2 "github.com/0xfbravo/brla/model"
)

type UseCases interface {
	// BuyStaticPix creates a new static pix on BRLA
	// See more: https://brla-superuser-api.readme.io/reference/superuserbuystaticpix
	BuyStaticPix(options *model2.BuyStaticPixOptions) (*model2.Pix, error)

	// GetBalanceOf retrieves the balance of a wallet directly from the blockchain
	GetBalanceOf(options *model2.BalanceOfOptions) (*model2.BalanceOf, error)

	// GetKYCStatus retrieves KYC status from BRLA
	GetKYCStatus(taxId string) (*model2.KYCHistory, error)

	// GetUsedLimit retrieves user used limit from BRLA
	GetUsedLimit(taxId string) (*model2.Limits, error)

	// KycLevelOne creates a new KYC level one on BRLA
	// See more: https://brla-superuser-api.readme.io/reference/superuserkyclevelone
	KycLevelOne(options *model2.KycLevelOneOptions) (*model2.KycLevelOne, error)

	// KycLevelTwo creates a new KYC level two on BRLA
	// See more: https://brla-superuser-api.readme.io/reference/superuserkycleveltwo
	KycLevelTwo(options *model2.KycLevelTwoOptions) error

	// Login logs in a super user into BRLA
	// See more: https://brla-superuser-api.readme.io/reference/superuserlogin
	Login(email string, password string) (*model2.Session, error)
}
