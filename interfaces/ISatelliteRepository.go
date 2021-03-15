package interfaces

import "github.com/juansecardozo/quasarop/models"

type ISatelliteRepository interface {
	FindByName(name string) (models.SatelliteModel, error)
	UpdateByName(models.SatelliteModel) (models.SatelliteModel, error)
	FindAll() ([]models.SatelliteModel, error)
}
