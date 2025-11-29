package repositories

import (
	"github.com/fabriciolfj/rules-elegibility/data"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func ProviderCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) Save(data data.CustomerData) error {
	return r.db.Save(&data).Error
}
