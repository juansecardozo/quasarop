package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/juansecardozo/quasar/interfaces"
	"github.com/juansecardozo/quasar/models"
	"github.com/savaki/trilateration"
)

type TopSecretService struct {
	interfaces.ISatelliteRepository
}

func (service *TopSecretService) ResolveTransmitter(req *http.Request) (models.TransmitterModel, error) {
	var w models.WrapperModel

	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()

	_ = d.Decode(&w)

	if len(w.Satellites) != 3 {
		return models.TransmitterModel{}, errors.New("not enough data")
	}

	satellites := make([]models.SatelliteModel, 0)
	for _, message := range w.Satellites {
		satellite, _ := service.FindByName(strings.Title(message.Name))
		satellite.SetDistance(message.Distance)
		satellite.SetMessage(message.Message)
		satellites = append(satellites, satellite)
	}

	location, err := service.GetLocation(satellites)
	if err != nil {
		return models.TransmitterModel{}, err
	}

	message, err := service.GetMessage(satellites)
	if err != nil {
		return models.TransmitterModel{}, err
	}

	return models.TransmitterModel{Position: location, Message: message}, nil
}

func (service *TopSecretService) GetLocation(satellites []models.SatelliteModel) (models.PositionModel, error) {
	if len(satellites) != 3 {
		return models.PositionModel{}, errors.New("not enough data")
	}

	points := make([]trilateration.Point, 0)
	for _, satellite := range satellites {
		if satellite.GetDistance() == 0 {
			return models.PositionModel{X: satellite.X, Y: satellite.Y}, nil
		}

		points = append(points, trilateration.Point{X: satellite.X, Y: satellite.Y, Z: 0, R: satellite.GetDistance()})
	}

	solution, err := trilateration.Solve(points[0], points[1], points[2])

	if err != nil {
		return models.PositionModel{}, err
	}

	s := solution.First()

	return models.PositionModel{X: s.X, Y: s.Y}, nil
}

func (service *TopSecretService) GetMessage(satellites []models.SatelliteModel) (string, error) {
	for pos, satellite := range satellites {
		if pos > 0 {
			if len(satellite.GetMessage()) != len((satellites[pos-1]).GetMessage()) {
				return "", errors.New("not enough data")
			}
		}
	}

	message := make([]string, len(satellites[0].GetMessage()))
	for i := 0; i < len(satellites[0].GetMessage()); i++ {
		for j := 0; j < len(satellites); j++ {
			text := satellites[j].GetMessage()[i]
			if text != "" {
				message[i] = text
				break
			}
			if j+1 == len(satellites) {
				return "", errors.New("not enough data")
			}
		}
	}

	return strings.Join(message, " "), nil
}
