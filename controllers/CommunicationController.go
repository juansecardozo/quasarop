package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/juansecardozo/quasar/interfaces"
	"github.com/juansecardozo/quasar/viewmodels"
)

type CommunicationController struct {
	interfaces.ICommunicationService
}

func (controller *CommunicationController) GetSatelliteResponse(res http.ResponseWriter, req *http.Request) {
	name := chi.URLParam(req, "name")

	satellite, err := controller.GetSatellite(name)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Header().Add("Content-Type", "application/json")
		json.NewEncoder(res).Encode(viewmodels.ErrorVM{Status: "error", Message: err.Error()})
	} else {
		res.Header().Add("Content-Type", "application/json")
		json.NewEncoder(res).Encode(viewmodels.SatelliteVM{Name: satellite.Name})
	}

}
