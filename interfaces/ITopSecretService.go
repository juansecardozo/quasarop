package interfaces

import (
	"net/http"

	"github.com/juansecardozo/quasar/models"
)

type ITopSecretService interface {
	ResolveTransmitter(request *http.Request) (models.TransmitterModel, error)
	GetLocation(satellites []models.SatelliteModel) (models.PositionModel, error)
	GetMessage(satellites []models.SatelliteModel) (string, error)
}
