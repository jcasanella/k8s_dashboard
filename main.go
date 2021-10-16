package main

import (
	"log"

	"github.com/jcasanella/k8s_dashboard/router"
)

func main() {
	router := router.Router()

	router.Static("/assets", "./assets")
	error := router.Run("localhost:8085")
	log.Fatal(error)
}
