package services

import (
	"github.com/google/uuid"
	"github.com/juansecardozo/quasar/interfaces"
	"github.com/juansecardozo/quasar/models"
)

type ChannelService struct {
	interfaces.IChannelRepository
}

func (service *ChannelService) StoreChannel() (models.ChannelModel, error) {
	id := uuid.New().String()
	channel, err := service.Store(id)

	if err != nil {
		return models.ChannelModel{}, err
	}

	return channel, nil
}
