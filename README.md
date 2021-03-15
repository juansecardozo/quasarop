# Operación Fuego de Quasar

Esta es una implementación de API REST con arquitectura limpia escrita en Go, con inyección de dependencias, siguiendo los principios SOLID.

## Pre-requisitos
Este proyecto requiere Go >= 1.15. Puede descargarse en el siguiente link:
[Downloads - The Go Programming Language (golang.org)](https://golang.org/dl/)

## Instalación

Clonar el repositorio

    git clone https://github.com/juansecardozo/quasarop
Instalar dependencias

    go get -u github.com/go-chi/chi/v5
    go get github.com/joho/godotenv
    go get github.com/lib/pq
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

## Arquitectura
La arquitectura está pensada para adaptarse al principio de separación de intereses, donde cada estructura debe tener una responsabilidad única para lograr un sistema desacoplado.

### controllers
Son los encargados de manejar todas las peticiones que llegan al «router».
El «struct» de un controlador implementa servicios a través de una interfaz. La implementación se inyecta en el momento de compilar.

### infrasctructures
Se encargan de la configuración para que el sistema se conecte a una fuente de datos externa. Sea una base de datos SQL, No-SQL o cualquier otro tipo de base de datos.

### interfaces
Son el puente entre diferentes dominios para que puedan interactuar entre sí, en este sistema, esta es la única forma de interactuar.

### models
Son representaciones de los objetos en la base de datos.

### services
Representan la capa de negocio. Es aquí donde se realiza la implementación de la lógica y se accede a la capa de datos.

### viewmodels
Son las representaciones de los recursos retornados como respuesta a una petición REST.

### servicecontainer
Es el encargado de inyectar las dependencias (implementaciones).