package services

import (
	"github.com/juansecardozo/quasarop/interfaces"
	"github.com/juansecardozo/quasarop/models"
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
