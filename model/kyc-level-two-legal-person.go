package model

type KycLevelTwoLegalPerson struct {
	KycLevelTwo
	PartnerCPF      string `json:"cPfPartner" validate:"required" example:"12345678901"`
	PartnerFullName string `json:"fullNamePartner" validate:"required" example:"John Doe"`
}
