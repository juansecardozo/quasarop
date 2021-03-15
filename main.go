package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	http.ListenAndServe(os.Getenv("DB_PORT"), ChiRouter().InitRouter())
}
