package dtos

import (
	"github.com/fabriciolfj/rules-elegibility/entities"
	"github.com/shopspring/decimal"
)

type CustomerRequest struct {
	Code     string          `json:"code"`
	Document string          `json:"document"`
	Income   decimal.Decimal `json:"income"`
}

type CustomerMessage struct {
	Code string `json:"code"`
}

func (dto CustomerRequest) ToEntity() *entities.Customer {
	return &entities.Customer{
		Code:     dto.Code,
		Document: dto.Document,
		Income:   dto.Income,
	}
}
