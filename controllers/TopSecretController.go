package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/juansecardozo/quasar/interfaces"
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
