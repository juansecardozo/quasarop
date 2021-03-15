package services

import (
	"github.com/juansecardozo/quasar/interfaces"
	"github.com/juansecardozo/quasar/models"
)

type CommunicationService struct {
	interfaces.ISatelliteRepository
}

func (service *CommunicationService) GetSatellite(name string) (models.SatelliteModel, error) {
	satellite, err := service.FindByName(name)

	if err != nil {
		return models.SatelliteModel{}, err
	}

	return satellite, nil
}
