package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/juansecardozo/quasarop/interfaces"
	"github.com/juansecardozo/quasarop/models"
	"github.com/juansecardozo/quasarop/viewmodels"
)

type TopSecretController struct {
	interfaces.ITopSecretService
}

func (controller *TopSecretController) ResolveTransmitterResponse(res http.ResponseWriter, req *http.Request) {

	transmitter, err := controller.ResolveTransmitter(req)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
	} else {
		res.Header().Add("Content-Type", "application/json")
		data, _ := json.Marshal(transmitter)
		res.Write(data)
	}
}

func (controller *TopSecretController) UpdateSatelliteResponse(res http.ResponseWriter, req *http.Request) {
	var s models.SatelliteModel

	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()

	_ = d.Decode(&s)

	s.Name = strings.Title(chi.URLParam(req, "satellite_name"))

	satellite, err := controller.UpdateSatellite(s)

	if err != nil {
		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(viewmodels.ErrorVM{Status: "error", Message: err.Error()})
	} else {
		res.Header().Add("Content-Type", "application/json")
		json.NewEncoder(res).Encode(viewmodels.SatelliteVM{Distance: satellite.Distance, Message: satellite.Message})
	}
}

func (controller *TopSecretController) ResolveSplitTransmitterResponse(res http.ResponseWriter, req *http.Request) {
	transmitter, err := controller.ResolveSplitTransmitter()

	if err != nil {
		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(viewmodels.ErrorVM{Status: "error", Message: err.Error()})
	} else {
		res.Header().Add("Content-Type", "application/json")
		data, _ := json.Marshal(transmitter)
		res.Write(data)
	}
}
