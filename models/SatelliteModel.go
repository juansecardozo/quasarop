package models

type SatelliteModel struct {
	Name     string
	X        float64
	Y        float64
	distance float64
	message  []string
}

func (model *SatelliteModel) SetDistance(distance float64) {
	model.distance = distance
}

func (model *SatelliteModel) GetDistance() float64 {
	return model.distance
}

func (model *SatelliteModel) SetMessage(message []string) {
	model.message = message
}

func (model *SatelliteModel) GetMessage() []string {
	return model.message
}
