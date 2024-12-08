package model

import (
	"github.com/0xfbravo/brla/enum"
	"time"
)

type KycLevelOneOptions struct {
	Birthdate  time.Time
	Document   string
	Name       string
	PersonType enum.PersonType
}
