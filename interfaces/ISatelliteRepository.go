package interfaces

import "github.com/juansecardozo/quasar/models"

type ISatelliteRepository interface {
	FindByName(name string) (models.SatelliteModel, error)
}
