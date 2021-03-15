package repositories

import (
	"github.com/juansecardozo/quasarop/interfaces"
	"github.com/juansecardozo/quasarop/models"
	"github.com/lib/pq"
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
	row.Scan(&satellite.Name, &satellite.X, &satellite.Y, (*pq.StringArray)(&satellite.Message), &satellite.Distance)

	return satellite, nil
}

func (repo *SatelliteRepository) UpdateByName(satellite models.SatelliteModel) (models.SatelliteModel, error) {
	err := repo.Execute("UPDATE satellites SET distance=$1, message=$2 WHERE name=$3", satellite.Distance, pq.Array(satellite.Message), satellite.Name)

	if err != nil {
		return models.SatelliteModel{}, err
	}

	return satellite, nil
}

func (repo *SatelliteRepository) FindAll() ([]models.SatelliteModel, error) {
	row, err := repo.Query("SELECT * FROM satellites")

	if err != nil {
		return []models.SatelliteModel{}, err
	}

	var satellite models.SatelliteModel

	satellites := make([]models.SatelliteModel, 0)
	for row.Next() {
		err := row.Scan(&satellite.Name, &satellite.X, &satellite.Y, (*pq.StringArray)(&satellite.Message), &satellite.Distance)
		if err != nil {
			return []models.SatelliteModel{}, err
		}
		satellites = append(satellites, satellite)
	}

	return satellites, nil
}
