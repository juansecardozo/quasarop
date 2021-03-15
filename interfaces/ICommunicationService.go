package interfaces

import "github.com/juansecardozo/quasar/models"

type ICommunicationService interface {
	GetSatellite(name string) (models.SatelliteModel, error)
}
