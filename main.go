package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	http.ListenAndServe(":8080", ChiRouter().InitRouter())
}
