package interfaces

import "github.com/juansecardozo/quasar/models"

type IChannelRepository interface {
	Store(id string) (models.ChannelModel, error)
}
