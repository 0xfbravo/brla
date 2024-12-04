package model

import (
	"github.com/0xfbravo/brla/enum"
)

type KycLevelOneOptions struct {
	Birthdate  string
	Document   string
	Name       string
	PersonType enum.PersonType
}
