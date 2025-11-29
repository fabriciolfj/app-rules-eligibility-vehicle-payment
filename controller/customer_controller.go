package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fabriciolfj/rules-elegibility/dtos"
	"github.com/fabriciolfj/rules-elegibility/usecases"
)

type CustomerController struct {
	useCaseSave *usecases.CustomerSaveUseCase
}

func ProviderCustomerController(uc *usecases.CustomerSaveUseCase) *CustomerController {
	return &CustomerController{
		useCaseSave: uc,
	}
}

func (controller *CustomerController) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var request dtos.CustomerRequest
	log.Printf("receive payload: %s", r.Body)

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		badRequest(w, err)
	}

	err = controller.useCaseSave.Execute(request.ToEntity())
	if err != nil {
		serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func serverError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func badRequest(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}
