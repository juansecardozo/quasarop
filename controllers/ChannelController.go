package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/juansecardozo/quasar/interfaces"
	"github.com/juansecardozo/quasar/viewmodels"
)

type ChannelController struct {
	interfaces.IChannelService
}

func (controller *ChannelController) StoreChannelResponse(res http.ResponseWriter, req *http.Request) {
	channel, err := controller.StoreChannel()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Header().Add("Content-Type", "application/json")
		json.NewEncoder(res).Encode(viewmodels.ErrorVM{Status: "error", Message: err.Error()})
	} else {
		res.Header().Add("Content-Type", "application/json")
		json.NewEncoder(res).Encode(viewmodels.ChannelVM{Id: channel.Id})
	}
}
