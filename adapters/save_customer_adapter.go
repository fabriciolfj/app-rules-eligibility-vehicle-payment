package adapters

import (
	"log"

	"github.com/fabriciolfj/rules-elegibility/data"
	"github.com/fabriciolfj/rules-elegibility/entities"
	"github.com/fabriciolfj/rules-elegibility/repositories"
)

type SaveCustomerAdapter struct {
	repository *repositories.CustomerRepository
}

func ProviderSaveCustomerAdapter(customerRepository *repositories.CustomerRepository) *SaveCustomerAdapter {
	return &SaveCustomerAdapter{
		repository: customerRepository,
	}
}

func (adapter *SaveCustomerAdapter) Process(customer *entities.Customer) error {
	customerData := data.ToData(customer)

	err := adapter.repository.Save(customerData)
	if err != nil {
		log.Fatal("failed to save customer", err)
	}

	return nil
}
