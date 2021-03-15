package main

import "github.com/go-chi/chi/v5"

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {
	communicationController := ServiceContainer().InjectCommunicationController()
	channelController := ServiceContainer().InjectChannelController()
	topSecretController := ServiceContainer().InjectTopSecretController()

	r := chi.NewRouter()
	r.Get("/satellites/{name}", communicationController.GetSatelliteResponse)
	r.Post("/channels", channelController.StoreChannelResponse)
	r.Post("/topsecret", topSecretController.ResolveTransmitterResponse)

	return r
}

var m *router

func ChiRouter() IChiRouter {
	return m
}
