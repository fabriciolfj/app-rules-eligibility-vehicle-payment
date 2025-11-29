package data

import (
	"time"

	"github.com/fabriciolfj/rules-elegibility/entities"
	"github.com/shopspring/decimal"
)

type CustomerData struct {
	Id          int64
	Code        string
	Document    string
	Income      decimal.Decimal
	DateCreated string
	DateUpdated string
}

func (data *CustomerData) ToEntity() entities.Customer {
	return entities.Customer{
		Code:     data.Code,
		Document: data.Document,
		Income:   data.Income,
	}
}

func ToData(entity *entities.Customer) CustomerData {
	return CustomerData{
		Document:    entity.Document,
		Code:        entity.Code,
		Income:      entity.Income,
		DateCreated: time.DateTime,
		DateUpdated: time.DateTime,
	}
}
