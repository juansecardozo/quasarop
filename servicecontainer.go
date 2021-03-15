package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/juansecardozo/quasarop/controllers"
	"github.com/juansecardozo/quasarop/infrastructures"
	"github.com/juansecardozo/quasarop/repositories"
	"github.com/juansecardozo/quasarop/services"
	_ "github.com/lib/pq"
)

type IServiceContainer interface {
	InjectTopSecretController() controllers.TopSecretController
}

type kernel struct{}

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
