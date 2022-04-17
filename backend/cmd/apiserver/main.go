package main

import (
	"backend/internal/app/apiserver"
	"log"
)

func main() {
	config := apiserver.NewConfig(":8000", "host=localhost dbname=phone sslmode=disable")

	if err := apiserver.Start(config); err != nil {
		log.Fatal()
	}
}
