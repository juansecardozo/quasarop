package main

import (
	"net/http"
	"os"
)

func main() {
	/* To be deleted */
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "quasar")
	/* To be deleted */
	http.ListenAndServe(":8080", ChiRouter().InitRouter())
}
