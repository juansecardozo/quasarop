package main

import "github.com/go-chi/chi/v5"

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {
	topSecretController := ServiceContainer().InjectTopSecretController()

	r := chi.NewRouter()
	r.Post("/topsecret", topSecretController.ResolveTransmitterResponse)
	r.Post("/topsecret_split/{satellite_name}", topSecretController.UpdateSatelliteResponse)
	r.Get("/topsecret_split", topSecretController.ResolveSplitTransmitterResponse)

	return r
}

var m *router

func ChiRouter() IChiRouter {
	return m
}
