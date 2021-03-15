package repositories

import (
	"github.com/juansecardozo/quasar/interfaces"
	"github.com/juansecardozo/quasar/models"
)

type ChannelRepository struct {
	interfaces.IDbHandler
}

func (repo *ChannelRepository) Store(id string) (models.ChannelModel, error) {
	err := repo.Execute("INSERT INTO channels (id) VALUES ($1)", id)

	if err != nil {
		return models.ChannelModel{}, nil
	}

	return models.ChannelModel{Id: id}, nil
}
