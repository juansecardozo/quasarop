package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/juansecardozo/quasar/controllers"
	"github.com/juansecardozo/quasar/infrastructures"
	"github.com/juansecardozo/quasar/repositories"
	"github.com/juansecardozo/quasar/services"
	_ "github.com/lib/pq"
)

type IServiceContainer interface {
	InjectCommunicationController() controllers.CommunicationController
	InjectChannelController() controllers.ChannelController
	InjectTopSecretController() controllers.TopSecretController
}

type kernel struct{}

func (k *kernel) InjectCommunicationController() controllers.CommunicationController {
	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	pgConn, err := sql.Open("postgres", strConn)

	if err != nil {
		log.Fatal(err)
	}

	pgHandler := &infrastructures.PostgresHandler{}
	pgHandler.Conn = pgConn

	satelliteRepository := &repositories.SatelliteRepository{pgHandler}
	communicationService := &services.CommunicationService{satelliteRepository}
	communicationController := controllers.CommunicationController{communicationService}

	return communicationController
}

func (k *kernel) InjectChannelController() controllers.ChannelController {
	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	pgConn, err := sql.Open("postgres", strConn)

	if err != nil {
		log.Fatal(err)
	}

	pgHandler := &infrastructures.PostgresHandler{}
	pgHandler.Conn = pgConn

	channelRepository := &repositories.ChannelRepository{pgHandler}
	channelService := &services.ChannelService{channelRepository}
	channelController := controllers.ChannelController{channelService}

	return channelController
}

func (k *kernel) InjectTopSecretController() controllers.TopSecretController {
	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	pgConn, err := sql.Open("postgres", strConn)

	if err != nil {
		log.Fatal(err)
	}

	pgHandler := &infrastructures.PostgresHandler{}
	pgHandler.Conn = pgConn

	satelliteRepository := &repositories.SatelliteRepository{pgHandler}
	TopSecretService := &services.TopSecretService{satelliteRepository}
	TopSecretController := controllers.TopSecretController{TopSecretService}

	return TopSecretController
}

var k *kernel

func ServiceContainer() IServiceContainer {
	return k
}
