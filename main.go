package main

import (
	"net/http"
	"os"
)

func main() {
	http.ListenAndServe(os.Getenv("DB_PORT"), ChiRouter().InitRouter())
}
