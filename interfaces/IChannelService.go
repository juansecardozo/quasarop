package interfaces

import "github.com/juansecardozo/quasar/models"

type IChannelService interface {
	StoreChannel() (models.ChannelModel, error)
}
