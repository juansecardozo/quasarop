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
		satellite.Distance = message.Distance
		satellite.Message = message.Message
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
		if satellite.Distance == 0 {
			return models.PositionModel{X: satellite.X, Y: satellite.Y}, nil
		}

		points = append(points, trilateration.Point{X: satellite.X, Y: satellite.Y, Z: 0, R: satellite.Distance})
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
			if len(satellite.Message) != len((satellites[pos-1]).Message) {
				return "", errors.New("not enough data")
			}
		}
	}

	message := make([]string, len(satellites[0].Message))
	for i := 0; i < len(satellites[0].Message); i++ {
		for j := 0; j < len(satellites); j++ {
			text := satellites[j].Message[i]
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

func (service *TopSecretService) UpdateSatellite(satellite models.SatelliteModel) (models.SatelliteModel, error) {
	satellite, err := service.UpdateByName(satellite)

	if err != nil {
		return models.SatelliteModel{}, err
	}

	return satellite, nil
}

func (service *TopSecretService) ResolveSplitTransmitter() (models.TransmitterModel, error) {
	satellites, err := service.FindAll()

	if err != nil {
		return models.TransmitterModel{}, err
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
