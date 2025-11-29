package entities

import "github.com/shopspring/decimal"

type Customer struct {
	Code     string
	Document string
	Income   decimal.Decimal
}
