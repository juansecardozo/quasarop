# Operación Fuego de Quasar

Esta es una implementación de API REST con arquitectura limpia escrita en Go, con inyección de dependencias junto con uso de «mocks», siguiendo los principios SOLID.

## Pre-requisitos
Este proyecto requiere Go > 1.14. Puede descargarse en el siguiente link:
[Downloads - The Go Programming Language (golang.org)](https://golang.org/dl/)

## Instalación

Clonar el repositorio

    git clone https://github.com/juansecardozo/quasarop
Instalar dependencias

    go get -u github.com/go-chi/chi/v5
    go get github.com/joho/godotenv
    go get github.com/lib/pq
    go get github.com/stretchr/testify
    go get github.com/vektra/mockery/v2/.../
    go get github.com/savaki/trilateration
   La aplicación hace uso de las siguientes variables de entrono:
	    
	   DB_HOST
	   DB_PORT
	   DB_USERNAME
	   DB_PASSWORD
	   DB_NAME
	   DB_SSL_MODE
	   PORT
> **Nota:** En el repositorio se encuentra un archivo .env.example haga una copia del mismo, renómbrelo a .env y modifique según requiera.

   Correr la aplicación

    go build && ./quasarop