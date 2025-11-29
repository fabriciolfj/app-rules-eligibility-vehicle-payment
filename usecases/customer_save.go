package usecases

import "github.com/fabriciolfj/rules-elegibility/entities"

type CustomerSaveGateway interface {
	Process(entity *entities.Customer) error
}

type NotifyCustomerCreatedGateway interface {
	Process(entity *entities.Customer)
}

type CustomerSaveUseCase struct {
	save   CustomerSaveGateway
	notify NotifyCustomerCreatedGateway
}

func ProviderCustomerSaveUseCase(gateway CustomerSaveGateway, notify NotifyCustomerCreatedGateway) *CustomerSaveUseCase {
	return &CustomerSaveUseCase{
		save:   gateway,
		notify: notify,
	}
}

func (uc CustomerSaveUseCase) Execute(entity *entities.Customer) error {
	err := uc.save.Process(entity)
	if err != nil {
		return err
	}

	uc.notify.Process(entity)
	return nil
}
