package model

import (
	"github.com/0xfbravo/brla/enum"
)

type KycLevelTwoOptions struct {
	Document        string
	PersonType      enum.PersonType
	PartnerCpf      *string
	PartnerFullName *string
	Files           *KycLevelTwo
}
