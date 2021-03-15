package repositories

import (
	"github.com/juansecardozo/quasar/interfaces"
	"github.com/juansecardozo/quasar/models"
)

type SatelliteRepository struct {
	interfaces.IDbHandler
}

func (repo *SatelliteRepository) FindByName(name string) (models.SatelliteModel, error) {
	row, err := repo.Query("SELECT * FROM satellites WHERE name = $1", name)

	if err != nil {
		return models.SatelliteModel{}, err
	}

	var satellite models.SatelliteModel

	row.Next()
	row.Scan(&satellite.Name, &satellite.X, &satellite.Y)

	return satellite, nil
}
