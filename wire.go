//go:build wireinject
// +build wireinject

package main

import (
	"github.com/fabriciolfj/rules-elegibility/adapters"
	"github.com/fabriciolfj/rules-elegibility/configuration"
	"github.com/fabriciolfj/rules-elegibility/controller"
	"github.com/fabriciolfj/rules-elegibility/listeners"
	"github.com/fabriciolfj/rules-elegibility/producers"
	"github.com/fabriciolfj/rules-elegibility/repositories"
	"github.com/fabriciolfj/rules-elegibility/usecases"
	"github.com/google/wire"
)

func InitController() (*controller.CustomerController, error) {
	wire.Build(
		configuration.ProviderDataBase,
		configuration.ProvideKafkaProperties,
		configuration.ProvideKafkaConfig,
		producers.ProviderRuleProcessProducer,
		repositories.ProviderCustomerRepository,
		adapters.ProviderSaveCustomerAdapter,
		adapters.ProviderNotifyCustomerCreatedAdapter,
		wire.Bind(new(usecases.CustomerSaveGateway), new(*adapters.SaveCustomerAdapter)),
		wire.Bind(new(usecases.NotifyCustomerCreatedGateway), new(*adapters.NotifyCustomerCreatedAdapter)),
		usecases.ProviderCustomerSaveUseCase,
		controller.ProviderCustomerController)

	return nil, nil
}

func InitListenerProcessWire() (*listeners.CustomerCreatedListener, error) {
	wire.Build(
		configuration.ProvideKafkaProperties,
		configuration.ProvideKafkaConfig,
		listeners.ProviderListenerCustomerCreated,
	)

	return nil, nil
}
